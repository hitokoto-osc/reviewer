package poll

import (
	"context"
	"fmt"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

func CalcReviewerAdoptionRate(ctx context.Context) error {
	g.Log().Debug(ctx, "开始计算审核员采纳率……")
	defer g.Log().Debug(ctx, "计算审核员采纳率完成！")
	// 暴力获取存在记录的审核员和管理员（且当前存在权限的）
	records, err := g.DB().Ctx(ctx).Raw(
		fmt.Sprintf("SELECT `%s` FROM `%s` WHERE `%s` in ("+
			"SELECT `%s` FROM `%s` WHERE `%s` = 1 OR `%s` = 1"+
			")",
			dao.PollUsers.Columns().UserId,
			dao.PollUsers.Table(),
			dao.PollUsers.Columns().UserId,
			dao.Users.Columns().Id,
			dao.Users.Table(),
			dao.Users.Columns().IsAdmin,
			dao.Users.Columns().IsReviewer,
		),
	).Array()
	if err != nil {
		return gerror.Wrap(err, "获取需要计算的用户失败")
	}
	userIDs := gconv.Ints(records)
	for _, userID := range userIDs {
		g.Log().Debugf(ctx, "开始获取用户 %d 的采纳率……", userID)
		adoptions, e := getUserAdoptions(ctx, gconv.Uint(userID))
		if e != nil {
			e = gerror.Wrapf(e, "获取用户 %d 采纳率失败：", userID)
			g.Log().Error(ctx, e)
			continue
		}
		adoptionsRate := float64(adoptions.Adoptions) / float64(adoptions.Total)
		g.Log().Debugf(ctx, "用户 %d 采纳率为 %f", userID, adoptionsRate)
		// 更新用户采纳率
		_, e = dao.PollUsers.Ctx(ctx).
			Where(dao.PollUsers.Columns().UserId, userID).
			Update(do.PollUsers{
				AdoptionRate: gconv.Float64(adoptionsRate),
			})
		if e != nil {
			e = gerror.Wrapf(e, "更新用户 %d 采纳率失败：", userID)
			g.Log().Error(ctx, e)
		}

		go func(userID int) {
			_, e = g.DB().GetCache().Remove(ctx, "user:poll:uid:"+gconv.String(userID))
			if e != nil {
				g.Log().Errorf(ctx, "清除用户 %d 缓存失败：%s", userID, e.Error())
			}
		}(userID)
	}
	g.Log().Debug(ctx, "计算审核员采纳率完成！")
	return nil
}

type UserAdoptions struct {
	Total     int `json:"total"`     // 总数
	Adoptions int `json:"adoptions"` // 采纳数
}

func getUserAdoptions(ctx context.Context, userID uint) (*UserAdoptions, error) {
	sql := fmt.Sprintf(`SELECT
		COUNT( 1 ) AS total,
			SUM(
				CASE
		WHEN `+"`%s`"+` = "超时自动处理" THEN 1 # 就当他通过了吧
		WHEN ( pipeline.`+"`%s`"+` = %d AND log.`+"`%s`"+` = %d ) THEN 1
		WHEN ( pipeline.`+"`%s`"+` = %d AND log.`+"`%s`"+` = %d ) THEN 1
		WHEN ( pipeline.`+"`%s`"+` = %d AND log.`+"`%s`"+` = %d ) THEN 1 ELSE 0
		END
		) AS adoptions
		FROM
		`+"`%s`"+` log
		JOIN `+"`%s`"+` pipeline ON log.%s = pipeline.%s
		WHERE
		log.user_id = %d`,
		dao.PollPipeline.Columns().Mark,
		// WHEN 批准
		dao.PollPipeline.Columns().Operate,
		int(consts.PollStatusApproved),
		dao.PollLog.Columns().Type,
		int(consts.PollMethodApprove),
		// WHEN 驳回
		dao.PollPipeline.Columns().Operate,
		int(consts.PollStatusRejected),
		dao.PollLog.Columns().Type,
		int(consts.PollMethodReject),
		// WHEN 需要修改
		dao.PollPipeline.Columns().Operate,
		int(consts.PollStatusNeedModify),
		dao.PollLog.Columns().Type,
		int(consts.PollMethodNeedModify),
		// 表
		dao.PollLog.Table(),
		dao.PollPipeline.Table(),
		dao.PollLog.Columns().PollId,
		dao.PollPipeline.Columns().PollId,
		userID,
	)
	var result UserAdoptions
	err := g.DB().Ctx(ctx).Raw(sql).Scan(&result)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户采纳率失败")
	}
	return &result, nil
}
