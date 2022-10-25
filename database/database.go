package database

import "real-time-forum/sqldb"

func CreateDB() {
	// user table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "users" (
				"userID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
				"nickname" TEXT NOT NULL UNIQUE,
				"age" TEXT NOT NULL UNIQUE, 
				"gender" TEXT NOT NULL,
				"firstname" TEXT,
				"lastname" TEXT,
				"email" TEXT NOT NULL,
				"passwordhash" BLOB NOT NULL
				); `)

	// post table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "posts" ( 
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
				);`)

	// comments table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "comments" ( 
		"commentID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"postID" INTEGER NOT NULL,
		"authorID" INTEGER NOT NULL,
		"author" TEXT NOT NULL,
		"text" TEXT NOT NULL, 
		"creationDate" TIMESTAMP,
		FOREIGN KEY(postID)REFERENCES posts(postID),
		FOREIGN KEY(authorID)REFERENCES users(userID)
		);`)

	// category table TABLE NOT USED YET
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "category" (
				"postID" TEXT REFERENCES post(postID), 
				"golang" INTEGER,
				"javascript" INTEGER,
				"rust" INTEGER,
				"python" INTEGER
				);`)

	// websockets table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "websockets" (
		"postID" TEXT REFERENCES post(postID), 
		"golang" INTEGER,
		"javascript" INTEGER,
		"rust" INTEGER,
		"python" INTEGER
		);`)

	// privatemessaging table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "privatechat" (
		"postID" TEXT REFERENCES post(postID), 
		"golang" INTEGER,
		"huh" INTEGER,
		"rust" INTEGER,
		"python" INTEGER
		);`)

	// sessions table
	sqldb.DB.Exec(`CREATE TABLE IF NOT EXISTS "session" ( 
				"sessionID" STRING NOT NULL PRIMARY KEY, 
				"userID" INTEGER NOT NULL,
				FOREIGN KEY(userID)REFERENCES users(userID)
				);`)
}
