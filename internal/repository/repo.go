package repository

import (
	"context"
	"errors"
	"myproject/internal/models"

	"gorm.io/gorm"
)

type DBRepository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	SignUp(ctx context.Context, userData models.SignupUsers) (models.SignupUsers, error)
	FetchUserWithEmail(ctx context.Context, userLogin models.LoginDB) (models.LoginDB, error)
	FetchAllUsers(ctx context.Context) ([]models.FetchAllUsers, error)
	FetchUserByID(ctx context.Context, fetchByID models.FetchUserByID) (models.FetchAllUsers, error)
	UpdateUserByID(ctx context.Context, updateById models.UpdateUserByID) (models.SignupUsersResp, error)
	DeleteUserByID(ctx context.Context, id int) (models.SignupUsersResp, error)
}

func NewRepository(db *gorm.DB) (RepositoryInterface, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &DBRepository{
		DB: db,
	}, nil
}
