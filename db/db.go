package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // CouchDB driver
	"github.com/joho/godotenv"
)

// Global CouchDB client
var DB *kivik.Client
var DBName = "holidays"

func InitDB() *kivik.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")

	// CouchDB connection string
	dsn := fmt.Sprintf("http://%s:%s@%s:%s/", user, pass, host, port)

	// Initialize CouchDB client
	client, err := kivik.New("couch", dsn)
	if err != nil {
		log.Fatal("Failed to connect to CouchDB:", err)
	}

	// Check if the database exists
	exists, err := client.DBExists(context.TODO(), DBName)
	if err != nil {
		log.Fatal("Error checking database existence:", err)
	}

	// Create the database if it does not exist
	if !exists {
		err = client.CreateDB(context.TODO(), DBName)
		if err != nil {
			log.Fatal("Error creating database:", err)
		}
		fmt.Println("Database created:", DBName)
	} else {
		fmt.Println("Database already exists:", DBName)
	}

	// Assign global client and return database reference
	DB = client
	return client.DB(DBName)
}
