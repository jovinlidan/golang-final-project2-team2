package user_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project2-team2/resources/user_resources"
	"golang-final-project2-team2/services/user_services"
	"golang-final-project2-team2/utils/error_formats"
	"golang-final-project2-team2/utils/error_utils"
	"golang-final-project2-team2/utils/success_utils"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userReq user_resources.UserRegisterRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserRegister(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UserLogin(c *gin.Context) {
	var userReq user_resources.UserLoginRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserLogin(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var userReq user_resources.UserUpdateRequest
	userIdParam := c.Param("userId")
	userIdToken := c.MustGet("user_id")
	if userIdToken != userIdParam {
		c.JSON(error_formats.NoAuthorization().Status(), error_formats.NoAuthorization())
		return
	}
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserUpdate(userIdParam, &userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userIdToken := c.MustGet("user_id")

	err := user_services.UserService.UserDelete(userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Your account has been successfully deleted"))
}
