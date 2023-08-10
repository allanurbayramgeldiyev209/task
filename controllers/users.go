package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"task/config"
	"task/models"

	"task/helpers"

	"github.com/gin-gonic/gin"
)

type LessonResponse struct {
	Day     int             `json:"day"`
	Lessons []models.Lesson `json:"lessons"`
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

func LoginUser(c *gin.Context) {

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

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
		"phone_number":  user.PhoneNumber,
	})

}

func GetLessons(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	usrID, hasUser := c.Get("user_id")
	if !hasUser {
		helpers.HandleError(c, 400, "user_id is required")
		return
	}
	userID, ok := usrID.(string)
	if !ok {
		helpers.HandleError(c, 400, "user_id must be string")
		return
	}

	var responses []LessonResponse
	rowsDay, err := db.Query(context.Background(), "SELECT DISTINCT(day_number) FROM lessons WHERE user_id = $1 AND deleted_at IS NULL ORDER BY day_number ASC", userID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsDay.Close()

	for rowsDay.Next() {
		var response LessonResponse
		if err := rowsDay.Scan(&response.Day); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowsLesson, err := db.Query(context.Background(), "SELECT order_number,name FROM lessons WHERE user_id = $1 AND day_number = $2 AND deleted_at IS NULL ORDER BY order_number ASC", userID, response.Day)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsLesson.Close()

		for rowsLesson.Next() {
			var lesson models.Lesson
			var lessonName sql.NullString
			if err := rowsLesson.Scan(&lesson.OrderNumber, &lessonName); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if lessonName.String == "" {
				lesson.Name = "null"
			} else {
				lesson.Name = lessonName.String
			}
			response.Lessons = append(response.Lessons, lesson)
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"weekle_lessons": responses,
	})

}
