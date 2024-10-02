//go:build wireinject
// +build wireinject

package di

import (
	"next-gen-job-hunting/api/auth"
	"next-gen-job-hunting/api/jobpost"
	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	user_job_post "next-gen-job-hunting/api/user-job-post"

	"next-gen-job-hunting/config/database"

	"github.com/google/wire"
)

func InitialiseUserService() *user.UserService {
	wire.Build(
		database.NewDB,

		user.NewUserRepository,
		user.NewUserService,
	)

	return &user.UserService{}
}

func InitialiseTokenService() *token.TokenService {
	wire.Build(
		database.NewDB,

		user.NewUserRepository,
		user.NewUserService,

		token.NewTokenRepository,
		token.NewTokenService,
	)

	return &token.TokenService{}
}

func InitializeUserController() *user.UserController {
	wire.Build(
		database.NewDB,

		user.NewUserRepository,
		user.NewUserService,

		user.NewUserValidationService,
		user.NewUserController,
	)

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

func InitializeAuthController() *auth.AuthController {
	wire.Build(
		database.NewDB,

		user.NewUserRepository,
		user.NewUserService,

		token.NewTokenRepository,
		token.NewTokenService,

		auth.NewAuthService,
		auth.NewAuthValidationService,
		auth.NewAuthController,
	)
	return &auth.AuthController{}
}

func InitialiseJobPostController() *jobpost.JobPostController {
	wire.Build(
		database.NewDB,

		jobpost.NewJobPostRepository,
		user_job_post.NewUserJobPostRepository,
		user_job_post.NewUserJobPostService,
		user_job_post.NewUserJobPostValidationService,
		jobpost.NewJobPostService,
		jobpost.NewJobPostValidationService,
		jobpost.NewJobPostController,
	)
	return &jobpost.JobPostController{}
}
