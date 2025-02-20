package app

import (
	"cc/cmd/models"
	"cc/cmd/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func App() {
	// Use command-line argument for country code, default to "US"
	countryCode := "IN"
	if len(os.Args) >= 2 {
		countryCode = os.Args[1]
	}

	// Your Calendarific API key
	apiKey := "nOOmoEMK4YY8UNYpEswN4rprA0OTXFSu" // replace with your api key
	url := fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=%s&year=%d", apiKey, countryCode, time.Now().Year())

	// Make the HTTP request
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Check for non-200 response
	if res.StatusCode != 200 {
		panic("Calendarific API not available")
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response
	var calendarResponse models.CalendarResponse
	err = json.Unmarshal(body, &calendarResponse)
	if err != nil {
		panic(err)
	}

	// Print holidays information and insert them into the database
	fmt.Printf("Public Holidays in %s for %d:\n", countryCode, time.Now().Year())
	for _, holiday := range calendarResponse.Response.Holidays {
		// Print each holiday
		fmt.Printf("%s - %s (International: %v)\n", holiday.Name, holiday.Date.ISO, holiday.International)

		// Insert the holiday into the database
		err := repositories.InsertHoliday(holiday)
		if err != nil {
			log.Printf("Error inserting holiday '%s' into the database: %v", holiday.Name, err)
		}
	}

	fmt.Println("Holiday data has been inserted into the database")
}
