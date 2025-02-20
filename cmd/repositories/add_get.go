package repositories

import (
	"cc/cmd/models"
	"cc/db"
	"context"
	"fmt"
	"log"

	"github.com/go-kivik/kivik/v4"
)

// AddH inserts a new holiday into the CouchDB database
func AddH(holiday models.Holiday) (models.Holiday, error) {
	DB := db.InitDB()

	// Create a document with a generated ID (ISO date as _id)
	docID := holiday.Date.ISO
	holidayData := map[string]interface{}{
		"_id":           docID,
		"name":          holiday.Name,
		"iso_date":      holiday.Date.ISO,
		"international": holiday.International,
	}

	// Insert the document into CouchDB
	_, err := DB.Put(context.TODO(), docID, holidayData)
	if err != nil {
		return models.Holiday{}, fmt.Errorf("error inserting holiday: %w", err)
	}

	return holiday, nil
}

// GetH retrieves all holiday documents from CouchDB
func GetH() ([]models.Holiday, error) {
	DB := db.InitDB()

	// Use kivik.Params() to correctly set options
	rows := DB.AllDocs(context.TODO(), kivik.Params(map[string]interface{}{
		"include_docs": true,
	}))
	if rows.Err() != nil {
		return nil, fmt.Errorf("error retrieving data: %w", rows.Err())
	}

	var holidays []models.Holiday

	// Iterate through each document
	for rows.Next() {
		var doc map[string]interface{} // Store raw JSON document
		if err := rows.ScanDoc(&doc); err != nil {
			log.Println("Skipping document due to scan error:", err)
			continue
		}

		// Convert raw JSON into models.Holiday
		var holiday models.Holiday
		holiday.ID, _ = doc["_id"].(int) // Convert _id field
		holiday.Name, _ = doc["name"].(string)

		// Extract ISO date properly
		if dateField, ok := doc["date"].(map[string]interface{}); ok {
			if iso, exists := dateField["iso"]; exists {
				if isoStr, valid := iso.(string); valid {
					holiday.Date.ISO = isoStr
				}
			}
		}

		holiday.International, _ = doc["international"].(bool)

		holidays = append(holidays, holiday)
	}

	return holidays, nil
}
