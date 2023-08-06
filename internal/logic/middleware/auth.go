package middleware

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

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
		g.Log().Error(r.GetCtx(), gerror.Wrap(err, "校验用户 Token 时发生错误"))
		r.Response.Status = http.StatusInternalServerError
		return
	}
	if !flag {
		r.Response.Status = http.StatusUnauthorized
		return
	}
	// 注入用户信息到上下文
	user, err := service.User().GetUserByToken(r.GetCtx(), token)
	if err != nil {
		g.Log().Error(r.GetCtx(), gerror.Wrap(err, "获取用户信息时发生错误"))
	} else if user == nil {
		r.Response.Status = http.StatusUnauthorized
		return
	}
	userPoll, err := service.User().GetPollUserByUserID(r.GetCtx(), user.Id)
	if err != nil {
		g.Log().Error(r.GetCtx(), gerror.Wrap(err, "获取用户投票信息时发生错误"))
	} else if userPoll == nil {
		g.Log().Error(r.GetCtx(), gerror.New("获取用户投票信息时发生错误，用户投票信息为空"))
		r.Response.Status = http.StatusInternalServerError
		return
	}
	userPattern := &model.UserCtxSchema{
		Users:  *user,
		Poll:   *userPoll,
		Role:   service.User().MustGetRoleByUser(r.GetCtx(), user),
		Status: service.User().MustGetUserStatusByUser(r.GetCtx(), user),
	}
	service.BizCtx().SetUser(r.GetCtx(), userPattern)
	r.Middleware.Next()
}

// AuthorizationAdminV1 用于 v1 接口校验用户是否登录且是否具有管理员权限
func (s *sMiddleware) AuthorizationAdminV1(r *ghttp.Request) {
	r.Middleware.Next()
}
