package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open database
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// access database from handlers package, need to create a pointer to db
	databaseHandler := &handlers.Forum{DB: database}

	// Check all required tables exist in database, and create them if they don't
	for _, table := range []string{"users", "posts", "comments", "categories", "sessions"} {
		handlers.CheckTablesExist(database, table)
	}

	defer database.Close()

	// Start hosting web server
	fileServer := http.FileServer(http.Dir("./static"))               // serve content from the static directory
	http.Handle("/static/", http.StripPrefix("/static/", fileServer)) // redirect any requests to the root URL to the static directory
	http.HandleFunc("/", databaseHandler.Home)
	http.HandleFunc("/login", databaseHandler.LoginHandler)
	http.HandleFunc("/register", databaseHandler.RegistrationHandler)
	fmt.Println("Server started at http://localhost:9000.")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}
