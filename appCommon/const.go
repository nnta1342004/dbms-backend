package appCommon

import "os"

const (
	DbTypeUser = iota
	DbTypeImage
	DbTypeItem
	DbTypeItemImage
	DbTypeCart
	DbTypeUserForgotPassword
	DbTypeOrder
	DbTypeGroupItem
	DbTypeItemOrder
	DbTypeEvent
)
const (
	UserVerification     = "USER_VERIFICATION"
	UserVerificationTime = "USER_VERIFICATION_TIME"
	UserSessionId        = "USER_SESSION_ID"
)

const (
	PluginAws      = "aws"
	PluginGorm     = "mysql"
	PluginJwt      = "jwt"
	PluginMailer   = "mail"
	PluginRedis    = "redis"
	PluginLocker   = "locker"
	PluginRabbitMQ = "rabbitmq"
)
const (
	LockOrder = "LOCK_ORDER_"
)
const (
	TopicSendMailRecoveryPassword = "send_mail_recovery_password"
	TopicSendMailVerificationCode = "send_mail_verification_code"
)

var (
	Host     = "smtp.gmail.com"
	Port     = "587"
	Username = os.Getenv("username")
	Password = os.Getenv("password")
)

const (
	RedisCacheItem      = "REDIS_CACHE_ITEM_"
	RediscacheItemGroup = "REDIS_CACHE_ITEM_GROUP_"
)
const (
	ExpiryAccessToken  = 60 * 60 * 12 * 365
	ExpiryRefreshToken = 60 * 30 * 24 * 7
	CurrentUser        = "user"
)
