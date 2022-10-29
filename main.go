package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"real-time-forum/chat"
	"real-time-forum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func setUpRoutes() {
	http.HandleFunc("/", handlers.IndexHandler)
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
}

// Just used for testing
func fetchUserRecords(db *sql.DB) {
    record, err := db.Query("SELECT * FROM users")
    checkErr(err)
    defer record.Close()
    for record.Next() {
        var userID int
        var username string
        var email string
        var firstname string
        var lastname string
        var age int
        var gender string
        record.Scan(&userID, &username, &email, &firstname, &lastname, &age, gender)
        fmt.Printf("User: %d %s %s %s %s %d %s", userID, username, email, firstname, lastname, age, gender)
    }
}

func checkTablesExist(db *sql.DB, table string) {
    _, table_check := db.Query("select * from " + table + ";")
    if table_check != nil {
        fmt.Println("Error: " + table + " table doesn't exist in database.")

        if table == "users" {
            fmt.Println("Creating users table...")
            users_table := `CREATE TABLE IF NOT EXISTS users (
                "userID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
                "username" TEXT NOT NULL UNIQUE,
                "email" TEXT NOT NULL,
                "firstname" TEXT,
                "lastname" TEXT,
                "age" INTEGER NOT NULL, 
                "gender" TEXT NOT NULL
                );`
                // "passwordhash" BLOB NOT NULL

            users, err := db.Prepare(users_table)
            checkErr(err)
            users.Exec()
            }
                
        if table == "posts" {
            fmt.Println("Creating posts table...")
            posts_table := `CREATE TABLE IF NOT EXISTS posts (
                "postID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
                "authorID" INTEGER NOT NULL,
                "author" TEXT NOT NULL,
                "title" TEXT NOT NULL, 
                "text" TEXT NOT NULL, 
                "category1" TEXT NOT NULL,
                "category2" TEXT NOT NULL,
                "category3" TEXT NOT NULL,
                "category4" TEXT NOT NULL,
                "creationDate" TIMESTAMP,
                FOREIGN KEY(authorID)REFERENCES users(userID)
                );`

            posts, err := db.Prepare(posts_table)
            checkErr(err)
            posts.Exec()
        }
            
        if table == "comments" {
            fmt.Println("Creating comments table...")
            comments_table := `CREATE TABLE IF NOT EXISTS comments (
                "commentID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
                "postID" INTEGER NOT NULL,
                "authorID" INTEGER NOT NULL,
                "author" TEXT NOT NULL,
                "text" TEXT NOT NULL, 
                "creationDate" TIMESTAMP,
                FOREIGN KEY(postID)REFERENCES posts(postID),
                FOREIGN KEY(authorID)REFERENCES users(userID)
                );`

            comments, err := db.Prepare(comments_table)
            checkErr(err)
            comments.Exec()
        }

            
        if table == "categories" {
            fmt.Println("Creating categories table...")
            categories_table := `CREATE TABLE IF NOT EXISTS categories (
                "postID" TEXT REFERENCES post(postID), 
                "golang" INTEGER,
                "javascript" INTEGER,
                "rust" INTEGER,
                "python" INTEGER
                );`

            categories, err := db.Prepare(categories_table)
            checkErr(err)
            categories.Exec()
        }
            
        if table == "sessions" {
            fmt.Println("Creating sessions table...")
            sessions_table := `CREATE TABLE IF NOT EXISTS sessions (
                "sessionID" STRING NOT NULL PRIMARY KEY, 
                "userID" INTEGER NOT NULL,
                FOREIGN KEY(userID)REFERENCES users(userID)
                );`

            sessions, err := db.Prepare(sessions_table)
            checkErr(err)
            sessions.Exec()
        }
    }
}

func main() {
    // Check if database exists
    if _, err := os.Stat("database.db"); os.IsNotExist(err) {
        fmt.Println("Database does not exist, creating...")
        // Create database
        db, err := os.Create("database.db")
        checkErr(err)
        db.Close()
        defer db.Close()
    } else {
        fmt.Println("Database exists, skipping creation.")
    }

    // Now we know the database exists, we can open it
	database, _ := sql.Open("sqlite3", "database.db")

	// Check all required tables exist in database, and create them if they don't
    for _, table := range []string{"users", "posts", "comments", "categories", "sessions"} {
        checkTablesExist(database, table)
    }
    fmt.Println("All tables exist in database.")
    defer database.Close()

    // Start hosting web server
    fileServer := http.FileServer(http.Dir("static")) // serve content from the static directory
    http.Handle("/static/", http.StripPrefix("/static/", fileServer))   // redirect any requests to the root URL to the static directory
    http.Handle("/", fileServer) 
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Server started at http://localhost:8080")
    

    // Insert user details to database
    // query, err := database.Prepare("INSERT INTO users(username, email, firstname, lastname, age, gender) values('username','email@example.com','firstname','lastname',20,'male')")
    // checkErr(err)
    // _, err = query.Exec()
    // checkErr(err)

	// fetchUserRecords(database)
}