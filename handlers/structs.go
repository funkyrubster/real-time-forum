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

var Warning struct {
	Warn string
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
