package handlers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver connects go with sql
)

type Forum struct {
	 *sql.DB
}


// func (forum *Forum) CreateUser(user User){
// 	// Insert user into database
// 	query, err1 := forum.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + user.Nickname + "','" + user.Email + "','" + user.Password+ "','" + user.Firstname+ "','" + user.Lastname + "'," + user.Age + ",'" + user.Gender + "')")
// 	if err1 != nil {
// 		log.Fatal(err1)
// 	}
// 	_, err1 = query.Exec()
// 			fmt.Println(err1)
// 			defer query.Close()
// }


// ------------------ check if the table exist if not, create one 

func CheckTablesExist(db *sql.DB, table string) {
	_, table_check := db.Query("select * from " + table + ";")
	if table_check != nil {
			fmt.Println("Error: " + table + " table doesn't exist in database.")

	if table == "users" {
			fmt.Println("Creating users table...")
			users_table := `CREATE TABLE IF NOT EXISTS users (
					"userID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
					"username" TEXT NOT NULL UNIQUE,
					"email" TEXT NOT NULL,
					"password" TEXT NOT NULL,
					"firstname" TEXT,
					"lastname" TEXT,
					"age" INTEGER NOT NULL, 
					"gender" TEXT NOT NULL
					CHECK (length("username") >= 3 AND length("username") <= 20)
					CHECK (("email") LIKE '%_@__%.__%')
					CHECK (length("password") >= 8)
					);`
					// "passwordhash" BLOB NOT NULL

			users, errUser := db.Prepare(users_table)
			if errUser != nil {
				log.Fatal(errUser)
			}
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

			posts, errTable := db.Prepare(posts_table)
			if errTable != nil {
				log.Fatal(errTable)
			}
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

			comments, errCommments := db.Prepare(comments_table)
			if errCommments != nil {
				log.Fatal(errCommments)
			}
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

			categories, errCategories := db.Prepare(categories_table)
			if errCategories != nil {
				log.Fatal(errCategories)
			}
			categories.Exec()
	}
			
	if table == "sessions" {
			fmt.Println("Creating sessions table...")
			sessions_table := `CREATE TABLE IF NOT EXISTS sessions (
					"sessionID" STRING NOT NULL PRIMARY KEY, 
					"userID" INTEGER NOT NULL,
					FOREIGN KEY(userID)REFERENCES users(userID)
					);`

			sessions, errSession := db.Prepare(sessions_table)
			if errSession != nil {
				log.Fatal(errSession)
			}
			sessions.Exec()
	}
	}
}

