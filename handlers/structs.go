package handlers

import (
	"time"
)

// var tpl *template.Template

type RegisterData struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"newusername"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"newpassword"`
}

type User struct {
	Username  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	userID  int
	session string
	max_age int
}

// 	// CreationDate time.Time
// }

// var user_session Cookie

// var CurrentUser User

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
