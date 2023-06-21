//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"
)

func BuildInjector() *Injector {
	wire.Build(
		// injector
		InjectorSet,
	)

	return new(Injector)
}
