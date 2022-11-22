package handlers

type RegisterData struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"newusername"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"newpassword"`
}

type UserProfile struct {
	User         []RegisterData
	CreatedPosts []Post
	Session      []UserSession
}

type Post struct {
	PostID    int
	UserID    int
	Username  string `json:"username"`
	Title     string
	Category  string
	Content   string `json:"postBody"`
	CreatedAt string
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

var Warning struct {
	Warn string
}

type Comment struct {
	CommentID int
	PostID    int
	UserId    string
	Content   string
	CreatedAt string
}
