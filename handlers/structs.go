package handlers


// type UserProfile struct {
// User LoginData 
// Post []PostFeed 
// Comments []Comment
// }

type RegisterData struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"newusername"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"newpassword"`
}


type LoginData struct {
	
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type PostFeed struct {
	PostID    int `json:"postid"`
	Username  string
	Title     string
	Content   string
	Category  string
	CreatedAt string
}


type Comment struct {
	CommentID int
	PostID    int
	UserId    string
	Content   string
	CreatedAt string
}
