package controllers

import (
	"context"
	"net/http"
	"task/config"
	"task/models"

	"task/helpers"

	"github.com/gin-gonic/gin"
)

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

func LoginCustomer(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request - den maglumatlar alynyar
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if !helpers.ValidatePhoneNumber(user.PhoneNumber) {
		helpers.HandleError(c, 400, "invalid phone number")
		return
	}

	// database - den telefon belgisi request - den gelen telefon belga den bolan maglumat cekilyar
	var id, password string
	if err := db.QueryRow(context.Background(), "SELECT id,password FROM users WHERE phone_number = $1 AND deleted_at IS NULL", user.PhoneNumber).Scan(&id, &password); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger request - den gelen telefon belgili user database - de yok bolsa onda error response edilyar
	if id == "" {
		helpers.HandleError(c, 404, "this user does not exist")
		return
	}

	// eger user bar bolsa onda paroly dogry yazypdyrmy yazmandyrmy sol barlanyar
	credentialError := helpers.CheckPassword(user.Password, password)
	if credentialError != nil {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	// maglumatlar dogry bolsa auth ucin access_toke bilen resfresh_token generate edilyar
	accessTokenString, refreshTokenString, err := helpers.GenerateAccessTokenForUser(user.PhoneNumber, id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// front tarapa ugratmak ucin admin - in id - si boyunca maglumatlary get edilyar
	adm, err := GetCustomerByID(id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
		"admin":         adm,
	})

}
