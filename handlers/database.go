package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver connects go with sql
)

type Forum struct {
	*sql.DB
}

// --------------- GET USER PROFILE -------------//

func (data *Forum) GetUserProfile(username string) UserProfile {

	user := UserProfile{}

	rows, err := data.DB.Query(`SELECT * FROM users where username= ?`, username)
	if err != nil {
		log.Fatal(err)
	}
	var userID int
	var firstname string
	var lastname string
	var email string
	var nickname string
	var password string
	var age int
	var gender string

	for rows.Next() {
		err := rows.Scan(&userID, &nickname, &email, &password, &firstname, &lastname, &age, &gender)
		if err != nil {
			log.Fatal(err)
		}

		user = UserProfile{
			User: User{
				Username:  nickname,
				Firstname: firstname,
				Lastname:  lastname,
				Email:     email,
			},
			CreatedPosts: data.GetPosts(username),
		}
	}
	return user
}

//----------------------- CREATE POST-------------------------//

func (data *Forum) CreatePost(post Post) {
	stmt, err := data.DB.Prepare("INSERT INTO posts (username, title, content, category, creationDate) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(post.Username, post.Title, post.Content, post.Category, post.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
}

// ------------------- GET POSTS -------------------//

func (data *Forum) GetPosts(username string) []Post {

	var posts []Post // slice of Post 



	rows, err := data.DB.Query(`SELECT * FROM posts WHERE username =?`, username)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
	  var post Post  // just a struct 
		err := rows.Scan(&post.PostID, &post.Username, &post.Title, &post.Content, &post.Category, &post.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		// post = Post{
		// 	// CreatedAt: post.CreatedAt,
		// 	// Comments: //get commnets func,
		// }
		posts = append(posts, post)
	}

	return posts
	

}

// ------------------CREATE SESSION	----------------------//

// InsertSession ...
func (data *Forum) InsertSession(sess UserSession) {
	stmnt, err := data.DB.Prepare("INSERT INTO sessions (cookieValue, userID, username) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("AddSession error inserting into DB: ", err)
	}
	defer stmnt.Close()
	stmnt.Exec(sess.session, sess.userID, sess.username)
}

// -------------- DELETE SESSION------------------	//

// User's cookie expires when browser is closed, delete the cookie from the database.
func (data *Forum) DeleteSession(w http.ResponseWriter, userID int) error {
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	stmt, err := data.DB.Prepare("DELETE FROM session WHERE userID=?;")
	// defer stmt.Close()
	stmt.Exec(userID)
	if err != nil {
		fmt.Println("DeleteSession err: ", err)
		return err
	}
	return nil
}

// -------------- GET SESSION ----------------//

func (data *Forum) GetSession() []UserSession {

	session := []UserSession{}

	rows, err := data.DB.Query(`SELECT * FROM sessions`)
	if err != nil {
		log.Fatal(err)
	}

	var userID int
	var cookieValue string
	var userName string

	for rows.Next() {
		err := rows.Scan(&userID, &cookieValue, &userName)
		if err != nil {
			log.Fatal(err)
		}
		session = append(session, UserSession{
			userID:   userID,
			session:  cookieValue,
			username: userName,
		})
	}
	return session
}

// ---------------- CREATE TABLES ----------------//

func CheckTablesExist(db *sql.DB, table string) {
	_, table_check := db.Query("select * from " + table + ";")
	if table_check != nil {
		fmt.Println("Error: " + table + " table doesn't exist in database.")

		if table == "users" {
			fmt.Println("Creating users table...")
			users_table := `CREATE TABLE IF NOT EXISTS users (
					"userID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
					"username" TEXT NOT NULL,
					"email" TEXT NOT NULL,
					"password" TEXT NOT NULL,
					"firstname" TEXT,
					"lastname" TEXT,
					"age" INTEGER NOT NULL, 
					"gender" TEXT NOT NULL
					);`

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
					"username" TEXT REFERENCES sesssion(userID),
					"title" TEXT NOT NULL, 
					"content" TEXT NOT NULL, 
					"category" TEXT NOT NULL,
					"creationDate" TIMESTAMP
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
				userID INTEGER NOT NULL,
				cookieValue TEXT NOT NULL UNIQUE,
				username TEXT REFERENCES users(username),
				FOREIGN KEY(userID) REFERENCES Users(userID)
					);`

			sessions, errSession := db.Prepare(sessions_table)
			if errSession != nil {
				log.Fatal(errSession)
			}
			sessions.Exec()
		}
	}
}

//--------------- CONNECT WITH MAIN.GO----------------//

func Connect(db *sql.DB) *Forum {
	// Check all required tables exist in database, and create them if they don't
	for _, table := range []string{"users", "posts", "comments", "categories", "sessions"} {
		CheckTablesExist(db, table)
	}
	return &Forum{
		DB: db,
	}
}
