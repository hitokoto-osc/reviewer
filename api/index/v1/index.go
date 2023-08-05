package v1

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/" tags:"Index" method:"get" summary:"获得项目信息"`
}

type GetInfoRes struct {
	Name      string      `json:"name" dc:"项目名称"`
	Version   string      `json:"version" dc:"版本号"`
	Donate    string      `json:"donate" dc:"捐赠信息"`
	Website   string      `json:"website" dc:"项目网站"`
	Feedback  g.MapStrStr `json:"feedback" dc:"反馈方式"`
	Copyright string      `json:"copyright" dc:"版权声明"`
}
