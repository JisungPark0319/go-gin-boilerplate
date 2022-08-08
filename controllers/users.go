package controllers

import (
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

var userModel = new(models.UserModel)

func (u UsersController) Login(c *gin.Context) {
	var loginForm forms.Login

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "jwt": token})
}

func (u UsersController) Register(c *gin.Context) {
	var registerForm forms.Register

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}