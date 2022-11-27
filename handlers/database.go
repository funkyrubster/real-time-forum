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

// Pulls specific user's data and posts data from database and returns it as a User struct
func (data *Forum) GetUserProfile(username string) UserProfile {
	// Used to store the user's profile information
	user := UserProfile{}

	// Get a specific user's information from the 'users' table
	rows, err := data.DB.Query(`SELECT * FROM users where username= ?`, username)
	if err != nil {
		log.Fatal(err)
	}

	// Used to store the user's data so we can add it to struct later on
	var userID int
	var firstname string
	var lastname string
	var email string
	var nickname string
	var password string
	var age int
	var gender string

	// Scans through each column in the 'users' row and stores the data in the variables above
	for rows.Next() {
		err := rows.Scan(&userID, &nickname, &email, &password, &firstname, &lastname, &age, &gender)
		if err != nil {
			log.Fatal(err)
		}

		// This contains the specific user's data as well as all of their posts
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

func (data *Forum) getLatestPosts() []Post {
	// Used to store all of the posts
	var posts []Post
	// Used to store invidiual post data
	var post Post

	rows, err := data.DB.Query(`SELECT * FROM posts`)
	if err != nil {
		log.Fatal(err)
	}

	// Scans through every post
	for rows.Next() {
	    // Populates post var with data from each post found in table
		err := rows.Scan(&post.PostID, &post.Username, &post.Content, &post.Hashtag, &post.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		// Adds each post found from specific user to posts slice
		posts = append(posts, post)
	}
	return posts
}

func (data *Forum) getLatestHashtags() []Hashtag {
	// Used to store all of the posts
	var hashtags []Hashtag
	// Used to store invidiual post data
	var hashtag Hashtag

	rows, err := data.DB.Query(`SELECT * FROM hashtags`)
	if err != nil {
		log.Fatal(err)
	}

	// Scans through every post
	for rows.Next() {
	    // Populates post var with data from each post found in table
		err := rows.Scan(&hashtag.ID, &hashtag.Name, &hashtag.Count)
		if err != nil {
			log.Fatal(err)
		}
		// Adds each post found from specific user to posts slice
		hashtags = append(hashtags, hashtag)
	}
	return hashtags
}

// Handles creation of new posts
func (data *Forum) CreatePost(post Post) {
	stmt, err := data.DB.Prepare("INSERT INTO posts (username, content, hashtag, creationDate) VALUES (?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	
	// Uses data from post variable to insert into posts table
	_, err = stmt.Exec(post.Username, post.Content, post.Hashtag, post.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
}

// // TODO: Needs to be fixed to return all hashtags and their counts
// func (data *Forum) GetHashtags(hashtag Hashtag) []Hashtag {
// 	var hashtags []Hashtag // slice of Post 

// 	rows, err := data.DB.Query(`SELECT * FROM hashtags`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 	  var hashtag Hashtag  // just a struct 
// 		err := rows.Scan(&hashtag.hashtagID, &hashtag.hashtagName, &hashtag.hashtagCount)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		hashtags = append(hashtags, hashtag)
// 	}
// 	// fmt.Println(hashtags)
// 	// fmt.Println("now returning hashtags")
// 	return hashtags
// }

// Pulls all posts from specific user and returns it as a slice of Post structs
func (data *Forum) GetPosts(username string) []Post {
	// Used to store all of the posts
	var posts []Post
	// Used to store invidiual post data
	var post Post

	rows, err := data.DB.Query(`SELECT * FROM posts WHERE username =?`, username)
	if err != nil {
		log.Fatal(err)
	}

	// Scans through every row where the username matches the username passed in
	for rows.Next() {
	    // Populates post var with data from each post found in table
		err := rows.Scan(&post.PostID, &post.Username, &post.Content, &post.Hashtag, &post.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		// Adds each post found from specific user to posts slice
		posts = append(posts, post)
	}
	return posts
}

// Pulls all hashtags and their counts, returns it as a slice of Hashtag structs
// func (data *Forum) GetHashtagData() []Hashtag {
// 	// Used to store all of the hashtags
// 	var hashtags []Hashtag
// 	// Used to store invidiual hashtag data
// 	var hashtag Hashtag

// 	rows, err := data.DB.Query(`SELECT * FROM hashtags`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Scans through each hashtag
// 	for rows.Next() {
// 	    // Populates hashtag var with data from each hashtag found in table
// 		err := rows.Scan(&hashtag.hashtagID, &hashtag.hashtagName, &hashtag.hashtagCount)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// Adds each hashtag found from specific user to hashtags slice
// 		hashtags = append(hashtags, hashtag)
// 	}
// 	fmt.Println("Hashtags just pulled from DB:\n", hashtags)
// 	return hashtags
// }

// Inserts session into sessions table
func (data *Forum) InsertSession(sess UserSession) {
	stmnt, err := data.DB.Prepare("INSERT INTO sessions (cookieValue, userID, username) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error inserting session into table:", err)
	}
	defer stmnt.Close()
	stmnt.Exec(sess.session, sess.userID, sess.username)
}

// TODO: Clarification
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

// Checks all sessions from sessions table and returns latest session
func (data *Forum) GetSession() []UserSession {
	// Used to store session data
	session := []UserSession{}

	// Checks all sessions from sessions table
	rows, err := data.DB.Query(`SELECT * FROM sessions`)
	if err != nil {
		log.Fatal(err)
	}

	// Used to store individual session data
	var userID int
	var cookieValue string
	var userName string

	// For each session found, populate the variable above
	for rows.Next() {
		err := rows.Scan(&userID, &cookieValue, &userName)
		if err != nil {
			log.Fatal(err)
		}

		// Overwrites every session, leaving only data for the latest session
		session = append(session, UserSession{
			userID:   userID,
			session:  cookieValue,
			username: userName,
		})
	}
	return session
}

// Used when starting server - Ensures all tables are created to avoid errors
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
					"content" TEXT NOT NULL, 
					"hashtag" TEXT NOT NULL,
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

		if table == "hashtags" {
			fmt.Println("Creating hashtags table...")
			hashtags_table := `CREATE TABLE IF NOT EXISTS hashtags (
				"hashtagID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				"hashtagName" TEXT NOT NULL,
				"hashtagCount" INTEGER NOT NULL
				);`

			hashtags, errHashtags := db.Prepare(hashtags_table)
			if errHashtags != nil {
				log.Fatal(errHashtags)
			}
			hashtags.Exec()

			fmt.Println("Inserting hashtags into hashtags table...")
			stmt, err := db.Prepare("INSERT INTO hashtags (hashtagName, hashtagCount) VALUES (?, ?);")
			if err != nil {
				log.Fatal(err)
			}

			// Used to store hashtag names
			hashtagSlice := make([]string, 7)
			hashtagSlice[0] = "Tech"
			hashtagSlice[1] = "Food"
			hashtagSlice[2] = "Art"
			hashtagSlice[3] = "Sports"
			hashtagSlice[4] = "Fitness"
			hashtagSlice[5] = "Travel"
			hashtagSlice[6] = "Misc"

			// insert all hashtags into hashtags table
			for _, hashtag := range hashtagSlice {
				_, err = stmt.Exec(hashtag, 0)
				if err != nil {
					log.Fatal(err)
				}
			}
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

// Check all required tables exist in database, and create them if they don't
func Connect(db *sql.DB) *Forum {
	for _, table := range []string{"users", "posts", "comments", "hashtags", "sessions"} {
		CheckTablesExist(db, table)
	}
	return &Forum{
		DB: db,
	}
}
