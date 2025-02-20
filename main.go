package main

import (
	"cc/app"
	"cc/cmd/handlers"
	"cc/db"

	"github.com/labstack/echo/v4"
)

func main() {

	// Connect to CouchDB
	db.InitDB()

	// Start app
	app.App()

	e := echo.New()

	//fetch and upload data to couchdb
	e.GET("/app", handlers.InD)

	//route directory
	e.GET("/", handlers.Home)

	//add / insert new holiday to db
	e.POST("/n", handlers.Add)

	//read all data --iso date is not read from data---(have to use "curl -X GET "http://admin:0000@localhost:5984/holidays/_all_docs?include_docs=true" )
	e.GET("/ga", handlers.GetA)

	//update holiday from db  --- (have to use user:pass before localhost to access as db is not accessable to public)
	e.PUT("/u/:id", handlers.Up)

	//delete specific holiday from db
	e.DELETE("/d/:iso_date", handlers.Del)

	//delete all holidays from db
	e.DELETE("/da", handlers.DelA)

	e.Logger.Fatal(e.Start(":8080"))
}
