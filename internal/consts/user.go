package consts

const (
	UserAccessTokenV1Length = 40 // 用户访问令牌长度
)

type UserRole string

const (
	RoleAdmin    UserRole = "管理员"
	RoleUser     UserRole = "普通用户"
	RoleReviewer UserRole = "审核员"
)

type UserStatus int

const (
	UserStatusSuspended UserStatus = -1 // 用户已被封禁
	UserStatusNormal    UserStatus = 1
)

// UserLoginType 用户登录类型
type UserLoginType string

const (
	UserLoginTypeUID      UserLoginType = "uid"      // UID 尝试
	UserLoginTypeEmail    UserLoginType = "email"    // Email 尝试
	UserLoginTypeToken    UserLoginType = "token"    // Token 尝试
	UserLoginTypeUsername UserLoginType = "username" // Username 尝试
)
