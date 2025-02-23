package handlers

import (
	"cc/cmd/models"
	"cc/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {

	// Create a new Holiday struct to bind the query params
	vac := models.Holiday{}

	// Extract query parameters from the URL
	vac.Name = c.QueryParam("Name")
	vac.Date.ISO = c.QueryParam("iso_date")
	vac.International = c.QueryParam("international") == "true" // Convert string to boolean

	// Validate the iso_date parameter
	if vac.Date.ISO == "" {
		return c.JSON(http.StatusBadRequest, "ISO date cannot be empty")
	}

	// Call the AddH function to insert the holiday
	nDate, err := repositories.AddH(vac)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, nDate)
}

func GetA(c echo.Context) error {

	//call the GetH function to get all holidays
	holidays, err := repositories.GetH()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, holidays)
}

func GetSH(c echo.Context) error {
	id := c.Param("id") // Get holiday ID from URL parameter

	// Fetch holiday data from the repository
	holiday, err := repositories.GetS(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	// Return holiday data as JSON response
	return c.JSON(http.StatusOK, holiday)
}
