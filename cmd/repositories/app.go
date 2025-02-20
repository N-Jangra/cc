package repositories

import (
	"cc/cmd/models"
	"cc/db"
	"context"
	"fmt"
)

// InsertHoliday inserts a holiday document into CouchDB
func InsertHoliday(holiday models.Holiday) error {
	DB := db.InitDB()

	// Convert Holiday struct into a document format
	doc := map[string]interface{}{
		"name":          holiday.Name,
		"iso_date":      holiday.Date.ISO,
		"international": holiday.International,
	}

	// Insert the document into CouchDB
	_, err := DB.Put(context.TODO(), holiday.Date.ISO, doc)
	if err != nil {
		return fmt.Errorf("failed to insert holiday: %v", err)
	}

	fmt.Printf("Successfully inserted holiday: %s\n", holiday.Name)
	return nil
}
