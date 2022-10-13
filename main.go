package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/handlers"
)

func main() {
	sqldb.ConnectDB()
	sqldb.CreateDB()

	path2 := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", path2))

	http.HandleFunc("/", handlers.IndexHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	sqldb.CloseDB()
}
