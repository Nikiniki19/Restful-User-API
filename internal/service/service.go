package service

import (
	"context"
	"myproject/internal/auth"
	"myproject/internal/cache"
	"myproject/internal/models"
	"myproject/internal/repository"
)

type Service struct {
	Repo repository.RepositoryInterface
	Auth auth.AuthInterface
	Rdb  cache.CacheInterface
}

type ServiceInterface interface {
	Signup(ctx context.Context, userDetails models.Users) (models.SignupUsersResp, error)
	Login(ctx context.Context, userLogin models.Login) (models.LoginResp, error)
	FetchAllUsers(ctx context.Context) ([]models.FetchAllUsers, error)
	FetchUserByID(ctx context.Context, fetchByID models.FetchUserByID) (models.FetchAllUsers, error)
	UpdateUserByID(ctx context.Context, updateById models.UpdateUserByID) (models.SignupUsersResp, error)
	DeleteUserByID(ctx context.Context, id int) (models.SignupUsersResp, error)
}

func NewService(repo repository.RepositoryInterface, auth auth.AuthInterface, rdb cache.CacheInterface) (ServiceInterface, error) {
	return &Service{
		Repo: repo,
		Auth: auth,
		Rdb:  rdb,
	}, nil
}
