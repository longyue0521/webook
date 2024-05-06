// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"github.com/ecodeclub/webook/internal/project"
	testioc "github.com/ecodeclub/webook/internal/test/ioc"
)

// Injectors from wire.go:

func InitModule() *project.Module {
	db := testioc.InitDB()
	module := project.InitModule(db)
	return module
}
