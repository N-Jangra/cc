package handlers

import (
	"cc/cmd/models"
	"cc/cmd/repositories"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler for UpdateH function
func Up(c echo.Context) error {
	id := c.Param("id") // Keep ID as string

	// Bind the request body to holiday struct
	user := models.Holiday{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Call UpdateH with a string ID
	updatedUser, err := repositories.UpdateH(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// handler for DeleteH function
func Del(c echo.Context) error {
	// Extract the iso_date from the URL parameter
	isoDate := c.Param("iso_date")
	if isoDate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ISO date cannot be empty"})
	}

	// Call the DeleteH function to delete the holiday from the database
	err := repositories.DeleteH(isoDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Holiday with ISO date %s has been deleted", isoDate)})
}

// handler for DellAll function
func DelA(c echo.Context) error {
	// Call the DelAll function to delete all holidays
	err := repositories.DelAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "All holidays deleted successfully"})
}
