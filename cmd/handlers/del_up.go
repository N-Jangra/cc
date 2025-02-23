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
	id := c.Param("id") // Extract the holiday ID from the URL

	var holiday models.Holiday
	if err := c.Bind(&holiday); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Manually map `iso_date` to `holiday.Date.ISO`
	if iso, ok := c.Get("iso_date").(string); ok {
		holiday.Date.ISO = iso
	}

	// Call the repository function to update the holiday
	updatedHoliday, err := repositories.UpdateH(holiday, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedHoliday)
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
