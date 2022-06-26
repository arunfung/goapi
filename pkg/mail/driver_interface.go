package mail

type Driver interface {
	// Send 发送验证码
	Send(email Email, config map[string]string) bool
}
