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

	http.HandleFunc("/test", handlers.TestHandler)

	http.HandleFunc("/", handlers.IndexHandler)      
	http.HandleFunc("/login", handlers.LoginHandler) 
	http.HandleFunc("/loginauth", handlers.LoginAuthHandler)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/registerauth", handlers.RegisterAuthHandler)
	http.HandleFunc("/addpost", handlers.AddPost)
	http.HandleFunc("/createpost", handlers.CreateAPost)
	http.HandleFunc("/filter", handlers.FilterHandle)

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
