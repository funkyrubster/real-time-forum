package main

import (
	"fmt"
	"log"
	"net/http"

	database "real-time-forum/database"
	"real-time-forum/handlers"
	sqldb "real-time-forum/sqldb"
)

func setUpRoutes() {
	http.HandleFunc("/", handlers.IndexHandler)
}

func main() {
	sqldb.ConnectDB()
	database.CreateDB()

	path2 := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", path2))

	setUpRoutes()
	// http.HandleFunc("/", handlers.IndexHandler)
	// http.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	sqldb.CloseDB()
}

// func main() {
// 	sqldb.ConnectDB()
// 	// sqldb.CreateDB() not creates yet

// 	path2 := http.FileServer(http.Dir("static"))

// 	http.Handle("/static/", http.StripPrefix("/static/", path2))

//

// 	fmt.Println("Server started at port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// 	sqldb.CloseDB()
// }
