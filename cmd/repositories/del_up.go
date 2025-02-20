package repositories

import (
	"cc/cmd/models"
	"cc/db"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-kivik/kivik/v4"
)

// UpdateH updates an existing holiday document
func UpdateH(holiday models.Holiday, id string) (models.Holiday, error) {
	DB := db.InitDB()

	// Retrieve the document's revision (_rev) before updating
	row := DB.Get(context.TODO(), id)
	if err := row.Err(); err != nil {
		if strings.Contains(err.Error(), "not_found") { // Check if error is due to missing document
			return models.Holiday{}, fmt.Errorf("holiday not found with ID: %s", id)
		}
		return models.Holiday{}, fmt.Errorf("error fetching holiday: %w", err)
	}

	var existing map[string]interface{}
	if err := row.ScanDoc(&existing); err != nil {
		return models.Holiday{}, fmt.Errorf("error reading existing document: %w", err)
	}

	// Ensure _id and _rev are preserved
	existing["name"] = holiday.Name
	existing["iso_date"] = holiday.Date.ISO
	existing["international"] = holiday.International

	// Save the updated document
	_, err := DB.Put(context.TODO(), id, existing)
	if err != nil {
		return models.Holiday{}, fmt.Errorf("error updating holiday: %w", err)
	}

	return holiday, nil
}

// DeleteH deletes a holiday document
func DeleteH(id string) error {
	DB := db.InitDB()

	// Retrieve document revision (_rev) before deletion
	row := DB.Get(context.TODO(), id)
	if err := row.Err(); err != nil { // Call row.Err() to check for errors
		return fmt.Errorf("holiday not found: %w", err)
	}

	var doc map[string]interface{}
	if err := row.ScanDoc(&doc); err != nil {
		return fmt.Errorf("error reading document: %w", err)
	}

	// Delete the document
	_, err := DB.Delete(context.TODO(), id, doc["_rev"].(string))
	if err != nil {
		return fmt.Errorf("error deleting holiday: %w", err)
	}

	return nil
}

func DelAll() error {
	DB := db.InitDB()

	// Use kivik.Params() to correctly set options
	rows := DB.AllDocs(context.TODO(), kivik.Params(map[string]interface{}{
		"include_docs": true,
	}))
	if rows.Err() != nil {
		return fmt.Errorf("error retrieving data")
	}

	// Slice to store bulk delete requests
	var bulkDocs []interface{}

	// Iterate through documents
	for rows.Next() {
		var doc map[string]interface{}
		if err := rows.ScanDoc(&doc); err != nil {
			log.Printf("Error scanning document: %v", err) // Log instead of skipping silently
			continue
		}

		// Ensure the document has _id and _rev
		if _, idOK := doc["_id"].(string); !idOK {
			log.Printf("Skipping invalid document (missing _id): %+v", doc)
			continue
		}
		if _, revOK := doc["_rev"].(string); !revOK {
			log.Printf("Skipping invalid document (missing _rev): %+v", doc)
			continue
		}

		// Mark document for deletion
		doc["_deleted"] = true

		// Add to bulk request
		bulkDocs = append(bulkDocs, doc)
	}

	// Perform bulk delete operation
	_, err := DB.BulkDocs(context.TODO(), bulkDocs)
	if err != nil {
		return fmt.Errorf("error deleting documents: %w", err)
	}

	fmt.Println("All holidays deleted successfully.")
	return nil
}
