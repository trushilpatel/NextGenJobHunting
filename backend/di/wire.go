//go:build wireinject
// +build wireinject

package di

import (
	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/user"

	"next-gen-job-hunting/config/database"

	"github.com/google/wire"
)

func InitializeUserController() *user.UserController {
	wire.Build(
		database.NewDB,
		user.NewUserRepository,
		user.NewUserService,
		user.NewUserController)

	return &user.UserController{}
}

func InitializeJobUrlController() *joburl.JobUrlController {
	wire.Build(
		database.NewDB,

		user.NewUserRepository,
		user.NewUserService,

		joburl.NewJobUrlRepository,
		joburl.NewJobUrlService,
		joburl.NewJobUrlController,
	)
	return &joburl.JobUrlController{}
}
