package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"./models"
)

func GetPage(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPages(db))
	}
}

func PutPage(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		//var page models.Page
		//// Map imcoming JSON body to the new Task
		//c.Bind(&page)
		//// Add a task using our new model
		//id, err := models.PutPage(db, page.Title, page.Body)
		//// Return a JSON response if successful
		//if err == nil {
		//	return c.JSON(http.StatusCreated, models.H{
		//		"created": id,
		//	})
		//	// Handle any errors
		//} else {
		//	return err
		//}
		return nil
	}
}
