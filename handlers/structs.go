package handlers

import "time"

/* ---------------------------------------------------------------- */
/*             USED FOR CREATING VARIABLES TO STORE DATA            */
/* ---------------------------------------------------------------- */

type RegisterData struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"newusername"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"newpassword"`
	LoggedIn  string
}

type User struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	LoggedIn  string
}

type OnlineActivity struct {
	Onlinestatus  string `json:"onlineuser"`
	OfflineStatus string `json:"offlineuser"`
}

type UserProfile struct {
	User         User
	CreatedPosts []Post
	Hashtags     []Hashtag
}

type Hashtags struct {
	Hashtags []Hashtag
}

type Post struct {
	PostID    int
	Username  string `json:"username"`
	Content   string `json:"postBody"`
	Hashtag   string
	CreatedAt time.Time
}

type Hashtag struct {
	ID    int
	Name  string `json:"name"`
	Count string `json:"count"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	LoggedIn string
}

type UserSession struct {
	username string
	userID   int
	session  string
	max_age  int
}

type Comment struct {
	CommentID int
	PostID    int
	Username  string `json:"username"`
	Content   string `json:"commentBody"`
	CreatedAt time.Time
}
