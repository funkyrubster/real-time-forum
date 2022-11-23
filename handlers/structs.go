package handlers

import "time"

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
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}


type UserProfile struct {
	User         User
	CreatedPosts []Post
}

type Post struct {
	PostID    int
	UserID    int
	Username  string `json:"username"`
	Content   string `json:"postBody"`
	CreatedAt time.Time
	Comments []Comment
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	username  string
	userID    int
	session   string
	max_age   int
}


type Comment struct {
	CommentID int
	PostID    int
	UserId    string
	Content   string
	CreatedAt string
}
