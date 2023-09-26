package poll

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/hitokoto-osc/reviewer/api/poll/adminV1"
)

func (c *ControllerAdminV1) UpdateOneMark(ctx context.Context, req *adminV1.UpdateOneMarkReq) (res *adminV1.UpdateOneMarkRes, err error) {
	mark, err := service.PollMarks().GetByID(ctx, int(req.ID))
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "获取标签失败")
	}
	if mark == nil {
		return nil, gerror.New("标签不存在")
	}
	err = copier.CopyWithOption(mark, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInvalidParameter, err, "参数错误")
	}
	err = service.PollMarks().Update(ctx, mark)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "修改标签失败")
	}
	return nil, nil
}
