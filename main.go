package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"real-time-forum/chat"
	database "real-time-forum/database"
	"real-time-forum/handlers"
	sqldb "real-time-forum/sqldb"
)

func setUpRoutes() {
	http.HandleFunc("/", handlers.IndexHandler)
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
}

func main() {
	sqldb.ConnectDB()
	database.CreateDB()

	path2 := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", path2)) //this handle the folder static

	setUpRoutes()
	// http.HandleFunc("/", handlers.IndexHandler)
	// http.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	sqldb.CloseDB()
}
