package poll

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/jinzhu/copier"

	"github.com/hitokoto-osc/reviewer/api/poll/adminV1"
)

func (c *ControllerAdminV1) AddOneMark(ctx context.Context, req *adminV1.AddOneMarkReq) (res *adminV1.AddOneMarkRes, err error) {
	mark := new(do.PollMark)
	err = copier.Copy(mark, req)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInvalidParameter, err, "参数错误")
	}
	err = service.PollMarks().Add(ctx, mark)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "添加投票标签失败")
	}
	return nil, nil
}
