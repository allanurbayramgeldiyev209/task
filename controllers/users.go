package controllers

import (
	"context"
	"net/http"
	"task/config"
	"task/models"

	"task/helpers"

	"github.com/gin-gonic/gin"
)

func LoginCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Hello",
	})
}
func RegisterUser(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	// request - den gelen maglumatlar alynyar
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.ValidateUser(user.PhoneNumber, "", true); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// parol hashlenyan
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// hemme zat yerbe yer bolsa maglumatlar users tablisa gosulyar
	_, err = db.Exec(context.Background(), "INSERT INTO users (phone_number,password) VALUES ($1,$2)", user.PhoneNumber, hashPassword)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"phone_number": user.PhoneNumber,
	})

}
