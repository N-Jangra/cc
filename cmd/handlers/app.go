package handlers

import (
	"cc/cmd/models"
	"cc/cmd/repositories"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func InD(c echo.Context) error {

	// Fetch data from the external API
	apiKey := "nOOmoEMK4YY8UNYpEswN4rprA0OTXFSu" // replace with your API key
	url := fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=IN&year=%d", apiKey, time.Now().Year())

	// make http request
	res, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data from external API"})
	}
	defer res.Body.Close()

	// Check for non-200 response
	if res.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "External API not available"})
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response body"})
	}

	// Parse the JSON response
	var CalendarResponse models.CalendarResponse
	if err := json.Unmarshal(body, &CalendarResponse); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse JSON response"})
	}

	// Insert holidays into the database
	for _, holiday := range CalendarResponse.Response.Holidays {
		if err := repositories.InsertHoliday(holiday); err != nil {
			fmt.Printf("Failed to insert holiday: %v\n", err)
		}
	}

	// Return success message if the insertion is successful
	return c.JSON(http.StatusOK, map[string]string{"message": "Holiday inserted successfully"})
}
