package main

import (
	"fmt"
	"forum/handlers"
	"forum/sqldb"
	"log"
	"net/http"
)

func main() {
	sqldb.ConnectDB()
	// sqldb.CreateDB() not creates yet

	path2 := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", path2))

	http.HandleFunc("/", handlers.IndexHandler) //not created yet

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	sqldb.CloseDB()
}
