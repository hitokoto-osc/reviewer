package hitokoto

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

type moveErrCollector struct {
	E    error
	UUID string
}

func (c *ControllerAdminV1) MoveList(ctx context.Context, req *adminV1.MoveListReq) (res *adminV1.MoveListRes, err error) {
	total := len(req.UUIDs)
	wg := &sync.WaitGroup{}
	wg.Add(total)

	failedChan := make(chan *moveErrCollector, total)
	for _, uuid := range req.UUIDs {
		go func(uuid string) {
			defer wg.Done()
			sentence, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, uuid)
			if err != nil {
				failedChan <- &moveErrCollector{
					E:    err,
					UUID: uuid,
				}
				return
			}
			err = service.Hitokoto().Move(ctx, sentence, req.Target)
			if err != nil {
				failedChan <- &moveErrCollector{
					E:    err,
					UUID: uuid,
				}
				return
			}
		}(uuid)
	}
	go func() {
		wg.Wait()
		close(failedChan)
	}()

	res = &adminV1.MoveListRes{
		IsSuccess:   true,
		Total:       total,
		FailedUUIDs: make([]string, 0),
		FailedDesc:  make(g.MapStrStr),
	}
	for failed := range failedChan {
		g.Log().Error(ctx, gerror.Stack(failed.E), failed.UUID)
		res.FailedUUIDs = append(res.FailedUUIDs, failed.UUID)
		res.FailedDesc[failed.UUID] = failed.E.Error()
	}
	if len(res.FailedUUIDs) > 0 {
		res.IsSuccess = false
	}
	return res, nil
}
