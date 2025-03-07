package service

import (
	"context"
	"errors"
	"fmt"
	"myproject/internal/models"
	"myproject/internal/utils"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Signup(ctx context.Context, userDetails models.Users) (models.SignupUsersResp, error) {
	PasswordOK := utils.ValidatePassword(userDetails.Password)
	if !PasswordOK {
		log.Error().Msg("password validation failed")
		return models.SignupUsersResp{}, errors.New("please provide proper password")
	}

	hashPwd, err := utils.HashPassword(userDetails.Password)
	if err != nil {
		log.Error().Err(errors.New("error in hashing the password"))
		return models.SignupUsersResp{}, errors.New("error in hashing the password")
	}

	userData := models.SignupUsers{
		UserName: userDetails.UserName,
		Email:    userDetails.Email,
		Password: hashPwd,
		Role:     userDetails.Role,
	}

	if userData.Role == "" {
		userData.Role = "user"
	}
	
	if userData.Role != "admin" && userData.Role != "user" {
		log.Error().Msg("Invalid role found in database")
		return models.SignupUsersResp{}, errors.New("invalid role")
	}
	users, err := s.Repo.SignUp(ctx, userData)
	if err != nil {
		log.Error().Err(errors.New("unable to signup"))
		return models.SignupUsersResp{}, errors.New("unable to signup")
	}
	data := models.SignupUsersResp{
		ID:      int(users.ID),
		Message: "user created",
	}
	return data, nil

}

func (s *Service) Login(ctx context.Context, userLogin models.Login) (models.LoginResp, error) {
	user := models.LoginDB{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	}

	login, err := s.Repo.FetchUserWithEmail(ctx, user)
	if err != nil {
		log.Error().Err(errors.New("unable to signin"))
		return models.LoginResp{}, errors.New("unable to sigin,please provide proper email and password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password))
	if err != nil {
		log.Error().Err(errors.New("unable to signin"))
		return models.LoginResp{}, errors.New("unable to sigin")
	}

	if login.Role != "admin" && login.Role != "user" {
		log.Error().Msg("Invalid role found in database")
		return models.LoginResp{}, errors.New("invalid role")
	}

	Id := strconv.Itoa(int(login.ID))
	if Id == "0" {
		log.Error().Msg("Invalid user ID retrieved")
		return models.LoginResp{}, errors.New("authentication error")
	}

	token, err := s.Auth.GenerateJWT(user.Email, login.Role)
	if err != nil {
		return models.LoginResp{}, err
	}

	err = s.Rdb.AddTokenToCache(ctx, Id, token)
	if err != nil {
		return models.LoginResp{}, err
	}
	cacheToken, err := s.Rdb.GetTokenFromCache(ctx, Id)
	fmt.Println("token from cache", cacheToken)
	if err == redis.Nil {
		cacheToken = token
	} else if err != nil {
		return models.LoginResp{}, err
	}

	return models.LoginResp{
		Message: "user logged in successfully",
		Token:   cacheToken,
	}, nil

}

func (s *Service) FetchAllUsers(ctx context.Context) ([]models.FetchAllUsers, error) {
	cachedUsers, err := s.Rdb.GetAllUsersFromCache(ctx, "all_users")
	fmt.Println("cached users are", cachedUsers)
	if err == nil && len(cachedUsers) > 0 {
		log.Info().Msg("Returning users from cache")
		return cachedUsers, nil
	}

	users, err := s.Repo.FetchAllUsers(ctx)
	if err != nil {
		log.Error().Err(errors.New("unable to fetch all the users"))
		return nil, errors.New("unable to fetch all the users")
	}

	err = s.Rdb.AddAllUsersToCache(ctx, users)
	if err != nil {
		return nil, err
	}

	return cachedUsers, nil
}

func (s *Service) FetchUserByID(ctx context.Context, fetchByID models.FetchUserByID) (models.FetchAllUsers, error) {
	fetchById, err := s.Repo.FetchUserByID(ctx, fetchByID)
	if err != nil {
		log.Error().Err(errors.New("ID is not present in db"))
		return models.FetchAllUsers{}, errors.New("ID is not present in db")
	}
	return fetchById, nil
}

func (s *Service) UpdateUserByID(ctx context.Context, updateById models.UpdateUserByID) (models.SignupUsersResp, error) {
	update, err := s.Repo.UpdateUserByID(ctx, updateById)
	if err != nil {
		log.Error().Err(errors.New("ID is not present in db"))
		return models.SignupUsersResp{}, errors.New("ID is not present in db")

	}
	return update, nil
}

func (s *Service) DeleteUserByID(ctx context.Context, id int) (models.SignupUsersResp, error) {
	delete, err := s.Repo.DeleteUserByID(ctx, id)
	if err != nil {
		log.Error().Err(errors.New("ID is not present in db"))
		return models.SignupUsersResp{}, errors.New("ID is not present in db")

	}
	return delete, nil
}
