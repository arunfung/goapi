package verifycode

import (
	"goapi/pkg/app"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// 增加前缀保持数据库整洁，出问题调试时也方便
				KeyPrefix: config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例：
//         verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(phone string) bool {

	// 生成验证码
	code := vc.generateVerifyCode(phone)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	// 发送短信
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
