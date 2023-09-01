package consts

const (
	UserAccessTokenV1Length = 40 // 用户访问令牌长度
)

// UserRole 用户角色，用于前端显示
type UserRole string

const (
	UserRoleGuest    UserRole = "guest"
	UserRoleAdmin    UserRole = "admin"
	UserRoleUser     UserRole = "user"
	UserRoleReviewer UserRole = "reviewer"
)

// UserRoleCode 用户角色代码
type UserRoleCode int

const (
	UserRoleCodeGuest    UserRoleCode = -1
	UserRoleCodeUser     UserRoleCode = 1
	UserRoleCodeReviewer UserRoleCode = 10
	UserRoleCodeAdmin    UserRoleCode = 1000
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

// UserPollPoints 用户可以的投票点数
type UserPollPoints int

const (
	UserPollPointsReviewer UserPollPoints = 1 // 审核员
	UserPollPointsNormal   UserPollPoints = 1 // 普通用户
	UserPollPointsAdmin    UserPollPoints = 2 // 管理员
)

type UserScoreType string

const (
	UserScoreTypeIncrement UserScoreType = "increment" // 增加
	UserScoreTypeDecrement UserScoreType = "decrement" // 减少
)
