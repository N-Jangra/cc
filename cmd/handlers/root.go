package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home serves as the welcome page
func Home(c echo.Context) error {
	response := map[string]string{
		"message": "Welcome to the Holiday API!",
		"routes":  "/n (Add), /ga (Get All), /g/:iso_date (Get by Date), /u/:id (Update), /d/:iso_date (Delete), /da (Delete All) /app (fetch data from api)",
		"status":  "API is running smoothly",
	}
	return c.JSON(http.StatusOK, response)
}
