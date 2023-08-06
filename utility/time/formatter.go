package time

import "github.com/gogf/gf/v2/os/gtime"

type Time gtime.Time

func (t Time) MarshalJSON() ([]byte, error) {
	time := gtime.Time(t)
	return []byte(`"` + time.Format("c") + `"`), nil
}
