package poll

import (
	"context"
	"math"
	"reflect"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

var fieldToStatus = map[string]consts.PollStatus{
	"Accept":     consts.PollStatusApproved,
	"Reject":     consts.PollStatusRejected,
	"NeedEdited": consts.PollStatusNeedModify,
}

var FieldAccept = "Accept"
var FieldReject = "Reject"
var FieldNeedEdited = "NeedEdited"

// DoPollRuling 处理投票
func DoPollRuling(ctx context.Context) error {
	g.Log().Debug(ctx, "开始裁决投票...")
	defer g.Log().Debug(ctx, "裁决投票任务执行完成")
	var (
		page      = 1
		pageSize  = 100
		totalPage int
	)
	query := dao.Poll.Ctx(ctx).
		Where(dao.Poll.Columns().Status, consts.PollStatusProcessing).
		WhereOr(dao.Poll.Columns().Status, consts.PollStatusOpen).
		WhereOr(dao.Poll.Columns().Status, consts.PollStatusOpenForCommonUser)
	total, err := query.Clone().Count()
	if err != nil {
		return gerror.Wrap(err, "获取投票数量失败")
	}
	if total == 0 {
		g.Log().Debug(ctx, "没有需要处理的投票")
		return nil
	}
	totalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	var polls []entity.Poll
	for {
		g.Log().Debugf(ctx, "正在处理第 %d 页，每页 %d 条记录，共 %d 页", page, pageSize, totalPage)
		err = query.Clone().Page(page, pageSize).Scan(&polls)
		if err != nil {
			return gerror.Wrap(err, "获取投票失败")
		}
		for i := 0; i < len(polls); i++ {
			poll := &polls[i]
			g.Log().Debugf(ctx, "正在处理投票 %d...", poll.Id)
			var maxField string
			// 比较 approve, reject, need_modify 三个字段，取最大值
			// 优先判断最大字段是否达到阈值，即优先级按数量大小排序
			if poll.Accept > poll.Reject {
				maxField = FieldAccept
			} else {
				maxField = FieldReject
			}
			if poll.NeedEdited > poll.Accept && poll.NeedEdited > poll.Reject {
				maxField = FieldNeedEdited
			}
			// 阈值策略
			totalTickets := poll.Accept + poll.Reject + poll.NeedEdited
			threshold := service.Poll().GetRulingThreshold(poll.IsExpandedPoll == 1, totalTickets)
			if reflect.ValueOf(poll).Elem().FieldByName(maxField).Int() < int64(threshold) {
				g.Log().Debugf(ctx, "投票 %d: 未达到阈值，跳过。", poll.Id)
				continue
			}
			e := service.Poll().DoRuling(ctx, poll, fieldToStatus[maxField])
			if e != nil {
				return gerror.Wrapf(e, "处理投票 %d 失败", poll.Id)
			}
		}
		page++
		if page > totalPage {
			break
		}
	}
	return nil
}
