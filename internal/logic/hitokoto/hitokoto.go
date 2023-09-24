package hitokoto

import (
	"context"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/samber/lo"

	"github.com/duke-git/lancet/v2/validator"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type sHitokoto struct{}

func init() {
	service.RegisterHitokoto(New())
}

func New() service.IHitokoto {
	return &sHitokoto{}
}

func (s *sHitokoto) GetHitokotoV1SchemaByUUID(ctx context.Context, uuid string) (*model.HitokotoV1WithPoll, error) {
	var (
		hitokoto *model.HitokotoV1WithPoll
		err      error
	)
	// 获取顺序 Cache -> Pending -> Sentence -> Refuse
	v, err := gcache.Get(ctx, "hitokoto:uuid:"+uuid)
	if err != nil {
		return nil, err
	} else if !v.IsNil() {
		return v.Val().(*model.HitokotoV1WithPoll), nil
	}
	// 从 Pending 中获取
	hitokotoInPending, err := s.GetPendingByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	} else if hitokotoInPending != nil {
		hitokoto, err = s.ConvertPendingToSchemaV1(ctx, hitokotoInPending)
		if err != nil {
			return nil, err
		}
		err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, consts.HitokotoV1SchemaCacheTime)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		return hitokoto, nil
	}
	// 从 Sentence 中获取
	hitokotoInSentence, err := s.GetSentenceByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if hitokotoInSentence != nil {
		hitokoto, err = s.ConvertSentenceToSchemaV1(ctx, hitokotoInSentence)
		if err != nil {
			return nil, err
		}
		err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, consts.HitokotoV1SchemaCacheTime)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		return hitokoto, nil
	}
	// 从 Refuse 中获取
	hitokotoInRefuse, err := s.GetRefuseByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	} else if hitokotoInRefuse == nil {
		return nil, nil
	}
	hitokoto, err = s.ConvertRefuseToSchemaV1(ctx, hitokotoInRefuse)
	if err != nil {
		return nil, err
	}
	err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, consts.HitokotoV1SchemaCacheTime)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return hitokoto, nil
}

var queryFields = []string{
	dao.Poll.Columns().Id,
	dao.Sentence.Columns().Uuid,
	dao.Sentence.Columns().Hitokoto,
	dao.Sentence.Columns().Type,
	dao.Sentence.Columns().From,
	dao.Sentence.Columns().FromWho,
	dao.Sentence.Columns().Creator,
	dao.Sentence.Columns().CreatorUid,
	dao.Sentence.Columns().Reviewer,
	dao.Sentence.Columns().CommitFrom,
	dao.Sentence.Columns().CreatedAt,
}

// GetList 获取句子列表
//
//nolint:gocyclo
func (s *sHitokoto) GetList(ctx context.Context, in *model.GetHitokotoV1SchemaListInput) (*model.GetHitokotoV1SchemaListOutput, error) {
	if in == nil {
		return nil, gerror.New("参数错误")
	}
	var query *gdb.Model

	// 判断是否需要获取所有状态的句子
	if in.Status == nil {
		query = g.DB().Ctx(ctx).Union(
			dao.Sentence.Ctx(ctx).Fields(append(queryFields, `"approved" AS status`)),
			dao.Pending.Ctx(ctx).Fields(append(queryFields, `"pending" AS status`)),
			dao.Refuse.Ctx(ctx).Fields(append(queryFields, `"refuse" AS status`)),
		)
	} else {
		switch *in.Status {
		case consts.HitokotoStatusApproved:
			query = dao.Sentence.Ctx(ctx).Fields(append(queryFields, `"approved" AS status`))
		case consts.HitokotoStatusPending:
			query = dao.Pending.Ctx(ctx).Fields(append(queryFields, `"pending" AS status`))
		case consts.HitokotoStatusRejected:
			query = dao.Refuse.Ctx(ctx).Fields(append(queryFields, `"rejected" AS status`))
		default:
			return nil, gerror.New("参数错误")
		}
	}

	if in.UUID != nil {
		query = query.Where(dao.Sentence.Columns().Uuid, *in.UUID)
		goto startQuery
	}

	query = query.Page(in.Page, in.PageSize)
	query = query.Order(`dao.Sentence.Columns().CreatedAt ` + in.Order)

	if in.Type != nil {
		query = query.Where(dao.Sentence.Columns().Type, *in.Type)
	}

	if in.Keywords != nil {
		keywords := strings.Split(*in.Keywords, " ")
		for _, keyword := range keywords {
			query = query.WhereOrLike(dao.Sentence.Columns().Hitokoto, "%"+keyword+"%")
		}
	}

	if in.From != nil {
		query = query.WhereOrLike(dao.Sentence.Columns().From, "%"+*in.From+"%")
	}

	if in.FromWho != nil {
		query = query.WhereOrLike(dao.Sentence.Columns().FromWho, "%"+*in.FromWho+"%")
	}

	if in.Creator != nil {
		query = query.WhereOrLike(dao.Sentence.Columns().Creator, "%"+*in.Creator+"%")
		if validator.IsInt(*in.Creator) {
			query = query.WhereOr(dao.Sentence.Columns().CreatorUid, *in.Creator)
		}
	}

startQuery:
	var results []model.HitokotoV1
	var total int
	err := query.ScanAndCount(&results, &total, false)
	if err != nil {
		return nil, gerror.Wrap(err, "查询句子失败")
	}
	items := make([]model.HitokotoV1WithPoll, len(results))
	// 生成需要查询的 UUID 列表
	uuids := lo.FilterMap(results, func(item model.HitokotoV1, index int) (string, bool) {
		// copy properties
		items[index] = model.HitokotoV1WithPoll{
			HitokotoV1: item,
		}
		if item.Status == consts.HitokotoStatusPending {
			return item.UUID, true
		}
		if item.Status == consts.HitokotoStatusApproved {
			items[index].PollStatus = consts.PollStatusApproved
		} else if item.Status == consts.HitokotoStatusRejected {
			items[index].PollStatus = consts.PollStatusRejected
		}
		return "", false // 不需要查询
	})
	// 查询投票状态
	polls, err := service.Poll().GetPollsBySentenceUUIDs(ctx, uuids)
	if err != nil {
		return nil, err
	}
	for i := range items {
		item := &items[i]
		if item.Status != consts.HitokotoStatusPending {
			continue // 跳过已经处理的句子
		}
		items[i].PollStatus = consts.PollStatusUnknown // 默认为未知
		for j := range polls {
			poll := &polls[j]
			if poll.SentenceUuid == item.UUID {
				item.PollStatus = consts.PollStatus(poll.Status)
				break
			}
		}
	}
	out := &model.GetHitokotoV1SchemaListOutput{
		Total:      total,
		Page:       in.Page,
		PageSize:   in.PageSize,
		Collection: items,
	}
	return out, nil
}

// Move 移动句子
//
//nolint:gocyclo
func (s *sHitokoto) Move(ctx context.Context, sentence *model.HitokotoV1WithPoll, target consts.HitokotoStatus) error {
	user := service.BizCtx().GetUser(ctx) // 获取用户
	if user == nil {
		return gerror.New("未登录")
	}

	if sentence == nil {
		return gerror.New("句子不存在")
	}

	var sentenceShouldBeRemoved *gdb.Model
	switch sentence.Status {
	case consts.HitokotoStatusApproved:
		sentenceShouldBeRemoved = dao.Sentence.Ctx(ctx).Where(dao.Sentence.Columns().Uuid, sentence.UUID)
	case consts.HitokotoStatusPending:
		break
	case consts.HitokotoStatusRejected:
		sentenceShouldBeRemoved = dao.Refuse.Ctx(ctx).Where(dao.Refuse.Columns().Uuid, sentence.UUID)
	default:
		return gerror.New("未知的句子状态")
	}
	// 移动句子
	if sentence.Status == consts.HitokotoStatusPending {
		// 只需要做裁决
		poll, e := service.Poll().GetPollBySentenceUUID(ctx, sentence.UUID)
		if e != nil {
			return gerror.Wrap(e, "获取投票失败")
		}
		e = service.Poll().DoRuling(ctx, poll, s.ConvertStatusToPollStatus(target))
		if e != nil {
			return gerror.Wrap(e, "处理投票失败")
		}
		return nil // 已经移动完成
	}

	var targetQuery *gdb.Model
	switch target {
	case consts.HitokotoStatusApproved:
		targetQuery = dao.Sentence.Ctx(ctx)
	case consts.HitokotoStatusPending:
		targetQuery = dao.Pending.Ctx(ctx)
	case consts.HitokotoStatusRejected:
		targetQuery = dao.Refuse.Ctx(ctx)
	default:
		return gerror.New("未知的句子目的状态")
	}

	e := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, e := sentenceShouldBeRemoved.TX(tx).Delete() // 删除原始句子
		if e != nil {
			return gerror.Wrap(e, "删除原始句子失败")
		} else if lo.Must(result.RowsAffected()) == 0 {
			return gerror.New("删除原始句子失败")
		}
		// 插入新的句子
		_, e = targetQuery.TX(tx).Unscoped().Insert(g.Map{
			dao.Sentence.Columns().Id:         sentence.ID,
			dao.Sentence.Columns().Uuid:       sentence.UUID,
			dao.Sentence.Columns().Hitokoto:   sentence.Hitokoto,
			dao.Sentence.Columns().Type:       sentence.Type,
			dao.Sentence.Columns().From:       sentence.From,
			dao.Sentence.Columns().FromWho:    sentence.FromWho,
			dao.Sentence.Columns().Creator:    sentence.Creator,
			dao.Sentence.Columns().CreatorUid: sentence.CreatorUID,
			dao.Sentence.Columns().Reviewer:   user.Id,
			dao.Sentence.Columns().CommitFrom: sentence.CommitFrom,
			dao.Sentence.Columns().CreatedAt:  sentence.CreatedAt,
		})
		if e != nil {
			return gerror.Wrap(e, "插入新的句子失败")
		}
		return nil
	})
	if e != nil {
		return e
	}
	service.Cache().ClearCacheAfterHitokotoUpdated(ctx, []string{sentence.UUID})
	return nil
}

// Delete 使用 UUIDs 删除句子
func (s *sHitokoto) Delete(ctx context.Context, uuids []string) error {
	if len(uuids) == 0 {
		return gerror.New("参数不能为空")
	}
	defer service.Cache().ClearCacheAfterHitokotoUpdated(ctx, uuids)
	return g.DB().Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		eg, egCtx := errgroup.WithContext(ctx)
		eg.Go(func() error {
			_, err := dao.Sentence.Ctx(egCtx).TX(tx).Delete(dao.Sentence.Columns().Uuid, uuids)
			return err
		})
		eg.Go(func() error {
			_, err := dao.Pending.Ctx(egCtx).TX(tx).Delete(dao.Pending.Columns().Uuid, uuids)
			return err
		})
		eg.Go(func() error {
			_, err := dao.Refuse.Ctx(egCtx).TX(tx).Delete(dao.Refuse.Columns().Uuid, uuids)
			return err
		})
		return eg.Wait()
	})
}

func (s *sHitokoto) UpdateByUUID(ctx context.Context, sentence *model.HitokotoV1WithPoll, do *model.DoHitokotoV1Update) error {
	var query *gdb.Model
	switch sentence.Status {
	case consts.HitokotoStatusApproved:
		query = dao.Sentence.Ctx(ctx).Where(dao.Sentence.Columns().Uuid, sentence.UUID).Unscoped()
	case consts.HitokotoStatusPending:
		query = dao.Pending.Ctx(ctx).Where(dao.Pending.Columns().Uuid, sentence.UUID).Unscoped()
	case consts.HitokotoStatusRejected:
		query = dao.Refuse.Ctx(ctx).Where(dao.Refuse.Columns().Uuid, sentence.UUID).Unscoped()
	}

	affectedRows, err := query.UpdateAndGetAffected(*do)
	if err != nil {
		return gerror.Wrap(err, "更新句子失败")
	}
	if affectedRows == 0 {
		return gerror.New("更新句子失败：句子不存在")
	}
	service.Cache().ClearCacheAfterHitokotoUpdated(ctx, []string{sentence.UUID})
	return nil
}
