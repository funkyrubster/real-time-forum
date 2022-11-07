package handlers

import (
	"time"
)

// var tpl *template.Template

// type User struct {
// 	UserID    int
// 	Nickname  string
// 	Age       string
// 	Gender    string
// 	Firstname string
// 	Lastname  string
// 	Email     string
// 	Password string
// 	// Username     string

// 	// CreationDate time.Time
// }

// var user_session Cookie

var CurrentUser User

var Warning struct {
	Warn string
}

// each session contains the username of the user and the time at which it expires
// type Session struct {
// 	UserID      int
// 	username    string
// 	sessionName string
// 	sessionUUID string
// 	expiry      time.Time
// }

type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}

type Pitem struct {
	PostID    int
	Nickname  string
	Title     string
	Text      string
	Category1 string
	Category2 string
	Category3 string
	Category4 string
	Comments  []Comm
}

type Category struct {
	CategoryID    int
	Catergoryname string
	PostID        Pitem
}

type Comm struct {
	CommentID int
	PostID    int
	Nickname  string
	Text      string
}
