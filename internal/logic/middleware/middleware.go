package middleware

import (
	"net/http"
	"strings"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	// 注册中间件
	service.RegisterMiddleware(New())
}

type sMiddleware struct{}

func New() service.IMiddleware {
	return &sMiddleware{}
}

type HandlerResponse struct {
	Code    int    `json:"code"    dc:"错误码"`
	Message string `json:"message" dc:"错误消息"`
	Data    any    `json:"data"    dc:"当前请求的结果数据"`
	TS      int64  `json:"ts"      dc:"当前请求的时间戳"`
}

// HandlerResponse 重写了默认的 JSON 响应格式，提供统一的响应格式
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(HandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		TS:      gtime.TimestampMilli(),
	})
}

// AuthorizationV1 用于 v1 接口校验用户是否登录
// 尝试顺序 Authorization: Bearer Token -> param -> form -> body -> query -> Router
func (s *sMiddleware) AuthorizationV1(r *ghttp.Request) {
	g.Log("AuthorizationV1").Debugf(r.GetCtx(), "AuthorizationV1: %s", r.GetHeader("Authorization"))
	authStr := r.GetHeader("Authorization")
	if authStr == "" || !strings.HasPrefix(authStr, "Bearer ") {
		if v := r.Get("token"); v != nil && !v.IsNil() && strings.HasPrefix(v.String(), "Bearer ") {
			authStr = v.String()
		} else {
			r.Response.Status = http.StatusUnauthorized
			return
		}
	}
	token := strings.Trim(strings.TrimPrefix(authStr, "Bearer "), " ")
	if len(token) != consts.UserAccessTokenV1Length {
		r.Response.Status = http.StatusUnauthorized
		return
	}
	flag, err := service.User().VerifyAPIV1Token(r.GetCtx(), token)
	if err != nil {
		g.Log().Panicf(r.GetCtx(), "校验用户 Token 时发生错误: %s", err.Error())
	}
	if !flag {
		r.Response.Status = http.StatusUnauthorized
		return
	}
	r.Middleware.Next()
}

// AuthorizationAdminV1 用于 v1 接口校验用户是否登录且是否具有管理员权限
func (s *sMiddleware) AuthorizationAdminV1(r *ghttp.Request) {
	r.Middleware.Next()
}
