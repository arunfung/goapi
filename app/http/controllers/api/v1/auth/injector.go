package auth

import (
	v1 "goapi/app/http/controllers/api/v1"

	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(
	wire.Struct(new(v1.BaseAPIController), "*"),
	wire.Struct(new(LoginController), "*"),
	wire.Struct(new(SignupController), "*"),
	wire.Struct(new(PasswordController), "*"),
	wire.Struct(new(VerifyCodeController), "*"),
	wire.Struct(new(Injector), "*"),
)

type Injector struct {
	LoginController      *LoginController
	SignupController     *SignupController
	PasswordController   *PasswordController
	VerifyCodeController *VerifyCodeController
}
