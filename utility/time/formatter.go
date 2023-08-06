package time

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type Time gtime.Time

func (t Time) MarshalJSON() ([]byte, error) {
	time := gtime.Time(t)
	return []byte(`"` + time.Format(consts.TimeFormat) + `"`), nil
}
