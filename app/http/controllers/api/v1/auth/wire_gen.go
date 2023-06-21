// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth

import (
	"goapi/app/http/controllers/api/v1"
)

// Injectors from wire.go:

func BuildInjector() *Injector {
	baseAPIController := v1.BaseAPIController{}
	loginController := &LoginController{
		BaseAPIController: baseAPIController,
	}
	signupController := &SignupController{
		BaseAPIController: baseAPIController,
	}
	passwordController := &PasswordController{
		BaseAPIController: baseAPIController,
	}
	verifyCodeController := &VerifyCodeController{
		BaseAPIController: baseAPIController,
	}
	injector := &Injector{
		LoginController:      loginController,
		SignupController:     signupController,
		PasswordController:   passwordController,
		VerifyCodeController: verifyCodeController,
	}
	return injector
}