package repository

import (
	"context"
	"errors"
	"myproject/internal/models"

	"github.com/rs/zerolog/log"
)

func (db *DBRepository) SignUp(ctx context.Context, userData models.SignupUsers) (models.SignupUsers, error) {
	if err := db.DB.Create(&userData).Error; err != nil {
		log.Error().Err(errors.New("failed to create the user"))
		return models.SignupUsers{}, err
	}
	return userData, nil

}

func (db *DBRepository) FetchUserWithEmail(ctx context.Context, userlogin models.LoginDB) (models.LoginDB, error) {
	if err := db.DB.WithContext(ctx).Table("signup_users").Select("password","id","role").Where("email=?", userlogin.Email).
		First(&userlogin).Error; err != nil {
		log.Error().Err(err).Msg("please provide valid email")
		return models.LoginDB{}, errors.New("please provide valid email")
	}
	return userlogin, nil
}

func (db *DBRepository) FetchAllUsers(ctx context.Context) ([]models.FetchAllUsers, error) {
	var results []models.FetchAllUsers
	if err := db.DB.WithContext(ctx).Table("signup_users").Select("*").Find(&results).
		Error; err != nil {
		log.Error().Err(err).Msg("please provide valid email")
		return nil, errors.New("please provide valid email")
	}
	return results, nil
}

func (db *DBRepository) FetchUserByID(ctx context.Context, fetchByID models.FetchUserByID) (models.FetchAllUsers, error) {
	var user models.FetchAllUsers
	if err := db.DB.WithContext(ctx).Table("signup_users").
		Select("*").Where("id=?", fetchByID.ID).
		First(&user).Error; err != nil {
		log.Error().Err(err).Msg("please provide valid ID")
		return models.FetchAllUsers{}, errors.New("please provide valid ID")
	}
	return user, nil
}

func (db *DBRepository) UpdateUserByID(ctx context.Context, updateById models.UpdateUserByID) (models.SignupUsersResp, error) {
	if err := db.DB.WithContext(ctx).Table("signup_users").
		Where("id=?", updateById.ID).
		Updates(
			map[string]interface{}{
				"email": updateById.Email,
			}).Error; err != nil {
		log.Error().Err(err).Msg("please provide valid ID")
		return models.SignupUsersResp{}, errors.New("please provide valid ID")
	}
	return models.SignupUsersResp{
		ID:      updateById.ID,
		Message: "user updated succesfully",
	}, nil
}

func (db *DBRepository) DeleteUserByID(ctx context.Context, id int) (models.SignupUsersResp, error) {
	var delete models.SignupUsers
	if err := db.DB.WithContext(ctx).Table("signup_users").Where("id=?", id).
		Delete(&delete).Error; err != nil {
		log.Error().Err(err).Msg("please provide valid ID")
		return models.SignupUsersResp{}, errors.New("please provide valid ID")
	}

	return models.SignupUsersResp{
		ID:      id,
		Message: "user deleted succesfully",
	}, nil
}
