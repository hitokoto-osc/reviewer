package common

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/" tags:"Index" method:"get" summary:"概览"`
}

type IndexRes struct {
	Name      string `json:"name" dc:"项目名称"`
	Version   string `json:"version" dc:"版本号"`
	Donate    string `json:"donate" dc:"捐赠信息"`
	Website   string `json:"website" dc:"项目网站"`
	Copyright string `json:"copyright" dc:"版权声明"`
}
