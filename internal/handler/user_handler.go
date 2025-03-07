package handler

import (
	"encoding/json"
	"myproject/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	var userDetails models.Users

	err := json.NewDecoder(c.Request.Body).Decode(&userDetails)
	if err != nil {
		log.Error().Err(err).Msg("Cannot decode the json")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please enter proper name,email and password",
		})
		return
	}
	validation := validator.New()
	err = validation.Struct(userDetails)
	if err != nil {
		log.Error().Err(err).Msg("validation failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, err := h.Ser.Signup(ctx, userDetails)
	if err != nil {
		log.Error().Err(err).Msg("Unable to signup the user")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("User signed up successfully")
	c.JSON(http.StatusOK, userId)

}

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var userLogin models.Login

	err := json.NewDecoder(c.Request.Body).Decode(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("Cannot decode the json")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please enter proper email and password",
		})
		return
	}

	validation := validator.New()
	err = validation.Struct(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("validation failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	login, err := h.Ser.Login(ctx, userLogin)
	if err != nil {
		log.Error().Err(err).Msg("Login failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("User logged in successfully")
	c.JSON(http.StatusOK, login)

}

func (h *Handler) FetchAllUsers(c *gin.Context) {
	ctx := c.Request.Context()

	fetchedUsers, err := h.Ser.FetchAllUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("fetch failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("Fetched All Users Successfully")
	c.JSON(http.StatusOK, fetchedUsers)
}

func (h *Handler) FetchUserByID(c *gin.Context) {
	ctx := c.Request.Context()
	var fetchByID models.FetchUserByID
	err := json.NewDecoder(c.Request.Body).Decode(&fetchByID)
	if err != nil {
		log.Error().Err(err).Msg("Cannot decode the json")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please enter proper email and password",
		})
		return
	}
	fetchUserByID, err := h.Ser.FetchUserByID(ctx, fetchByID)
	if err != nil {
		log.Error().Err(err).Msg("fetch failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("Fetched All Users Successfully")
	c.JSON(http.StatusOK, fetchUserByID)
}

func (h *Handler) UpdateUserByID(c *gin.Context) {
	ctx := c.Request.Context()
	var updateUser models.UpdateUserByID

	err := json.NewDecoder(c.Request.Body).Decode(&updateUser)
	if err != nil {
		log.Error().Err(err).Msg("Cannot decode the json")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please enter proper email and password",
		})
		return
	}
	updateID, err := h.Ser.UpdateUserByID(ctx, updateUser)
	if err != nil {
		log.Error().Err(err).Msg("update failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("user updated successfully")
	c.JSON(http.StatusOK, updateID)
}

func (h *Handler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	id1, _ := strconv.Atoi(id)
	deleteUser, err := h.Ser.DeleteUserByID(ctx, id1)
	if err != nil {
		log.Error().Err(err).Msg("deletion failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info().Msg("user updated successfully")
	c.JSON(http.StatusOK, deleteUser)

}
