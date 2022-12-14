package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Try to parse the index.html file and if it fails, log the error
func (data *Forum) Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, ""); err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}

}

// Handles receiving the comment data and adding it to the 'comments' table in the database
func (data *Forum) Comment(w http.ResponseWriter, r *http.Request) {

	var comment Comment

	// Decode the JSON data from the request body into the comment variable
	json.NewDecoder(r.Body).Decode(&comment)

	w.Write([]byte("ok"))

	// feches current session value
	x, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal(err)
	}
	sessionvalue := x.Value

	sess := data.GetSession(sessionvalue)
	time := time.Now()

	data.CreateComment(Comment{
		PostID:    comment.PostID,
		Username:  sess.username,
		Content:   comment.Content,
		CreatedAt: time,
	})
}

func (data *Forum) SendComments(w http.ResponseWriter, r *http.Request){

var comment Comment

json.NewDecoder(r.Body).Decode(&comment.PostID)


fmt.Println(comment.PostID)

}




// Handles receiving the post data and adding it to the 'posts' table in the database
func (data *Forum) Post(w http.ResponseWriter, r *http.Request) {
	// Decodes posts data into post variable
	var post Post

	// Decode the JSON data from the request body into the post variable
	json.NewDecoder(r.Body).Decode(&post)

	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	// feches current session value
	x, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal()
	}
	sessionvalue := x.Value

	// Convert data into variables for easier use
	hashtag := post.Hashtag
	time := time.Now()
	content := post.Content

	sess := data.GetSession(sessionvalue)

	// Inserts post into the 'posts' table of the database
	data.CreatePost(Post{
		//username from current session
		Username:  sess.username,
		Content:   content,
		Hashtag:   hashtag,
		CreatedAt: time,
	})

}


func (data *Forum) SendLatestPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sendLatestPosts() called")

	// Send user information back to client using JSON format
	posts := data.getLatestPosts()
	// fmt.Println(posts)
	// fmt.Println(userInfo)
	js, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
	w.Write([]byte(js))

	fmt.Println("sendLatestPosts() sent to JS")
}

// Updates hashtag count for specific hashtag when called
func (data *Forum) UpdateHashtag(w http.ResponseWriter, r *http.Request) {
	// Decodes posts data into post variable
	var hashtag Hashtag

	// Decode the JSON data from the request body into the post variable
	json.NewDecoder(r.Body).Decode(&hashtag)

	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	// Convert data into variables for easier use
	hashID := hashtag.ID
	hashName := hashtag.Name
	hashCount := hashtag.Count

	fmt.Println("hashID:", hashID)
	fmt.Println("hashName:", hashName)
	fmt.Println("hashtagCount:", hashCount)

	// Updates hashtag count in the 'hashtags' table of the database
	data.UpdateHashtagCount(Hashtag{
		ID:    hashID,
		Name:  hashName,
		Count: hashCount,
	})

}

func (data *Forum) SendLatestHashtags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SendLatestHashtags() called")

	// Send user information back to client using JSON format
	hashtags := data.getLatestHashtags()
	// fmt.Println(userInfo)
	js, err := json.Marshal(hashtags)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
	w.Write([]byte(js))

	fmt.Println("SendLatestHashtags() sent to JS")
}

// Handles the registration of new users - validates the data and adds it to the 'users' table in database
func (data *Forum) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Decodes registration data into user variable
	var user RegisterData
	json.NewDecoder(r.Body).Decode(&user)

	// Used in conjunction with the 'strings.ContainsAny' function to ensure the age entered is strictly numeric
	numChars := "0123456789"

	// Ensures all required fields are filled out, and that the age is strictly numeric
	if len(user.Firstname) == 0 || len(user.Lastname) == 0 || len(user.Email) == 0 || len(user.Username) == 0 || (len(user.Age) == 0 || !strings.ContainsAny(user.Age, numChars)) || user.Gender == "Gender" || len(user.Password) == 0 {
		// This HTTP status code is then checked in authentication.js and the user is alerted to the missing/invalid fields
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		// Uses web socket to read the information
		w.Header().Set("Content-type", "application/text")

		// These are initially false, and are only set to true if the email/username is not found in the database (registration is available and will not overwrite existing data)
		emailValid := false
		usernameValid := false
		user.LoggedIn = "false"

		/* ---------------------------------------------------------------- */
		/*       CHECKING IF EMAIL/USERNAME ALREADY EXISTS IN DATABASE      */
		/* ---------------------------------------------------------------- */

		/* --- Queries through each table, checks if data already exists -- */

		// EMAIL CHECK
		row := data.DB.QueryRow("select email from users where email= ?", user.Email)
		temp := "" // If email is not found, temp variable will remain empty
		row.Scan(&temp)
		if temp == "" {
			emailValid = true
		}

		// USERNAME CHECK
		row = data.DB.QueryRow("select username from users where username= ?", user.Username)
		temp = "" // If username is not found, temp variable will remain empty
		row.Scan(&temp)
		if temp == "" {
			usernameValid = true
		}

		// If both email and username are valid, we can successfully register the user into the database
		if emailValid && usernameValid {
			// Generates hash from password
			var passwordHash []byte
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Error hashing password:", err)
				return
			}

			// Inserts registration data into the 'users' table of the database
			query, err := data.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender,loggedin) VALUES(?, ?, ?, ?, ?, ?, ?, ?);")
			if err != nil {
				log.Fatal(err)
			}

			_, err = query.Exec(user.Username, user.Email, string(passwordHash), user.Firstname, user.Lastname, user.Age, user.Gender, user.LoggedIn)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("SUCCESS: User successfully registered into users table.")
			w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
		} else {
			fmt.Println("ERROR: Username or email already exists.")
			w.WriteHeader(http.StatusBadRequest) // Checked in authentication.js, alerts user
		}
	}
}

// Handles the login of existing users - validates the data and checks if it exists in the 'users' table in database
func (data *Forum) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Decodes session and login data into variables
	var sess UserSession
	var user LoginData

	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-type", "application/text")

	// Only set to true if the email/username IS found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	// Checks if user entered an email or username
	enteredEmail := strings.Contains(user.Username, "@")

	/* ---------------------------------------------------------------- */
	/*               CHECKING EMAIL/USER PASS COMBINATIONS              */
	/* ---------------------------------------------------------------- */

	/* --- Queries through each table, checks if data exists --- */

	// EMAIL CHECK
	if enteredEmail {
		// Checks if email/pass combination exists in database
		var passwordHash string
		row := data.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Username)
		err := row.Scan(&passwordHash)
		if err != nil {
			fmt.Println("Error with password hash:", err)
		}
		// If the password hash matches the password entered, the email/pass combination is valid
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err == nil {
			emailPassCombinationValid = true
		}
	} else {
		// Checks if username/pass combination exists in database
		var passwordHash string
		row := data.DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username)
		err := row.Scan(&passwordHash)
		if err != nil {
			fmt.Println("Error with password hash:", err)
		}
		// If the password hash matches the password entered, the user/pass combination is valid
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err == nil {
			userPassCombinationValid = true
		}
	}
	var usID int = 5

	// If either combination is valid, we can successfully log the user in
	if emailPassCombinationValid || userPassCombinationValid {
		fmt.Println("SUCCESS: User logged in.")

		row := data.DB.QueryRow("SELECT userID FROM users WHERE username = ?;", user.Username)
		err := row.Scan(&usID)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("usID:", usID)
		// fmt.Println("user.Username:", user.Username)

		// Creates a new session for the user
		sess.username = user.Username
		sess.userID = usID
		sess.max_age = 18000
		sess.session = (uuid.NewV4().String() + "&" + strconv.Itoa(sess.userID))
		user.LoggedIn = "true"

		// Set client cookie for "session_token" as session token we just generated, also set expiry time to 120 minutes
		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: sess.session,
			MaxAge: 900,
		})

		// x := sess.session + "&" + strconv.Itoa(sess.userID)
		// fmt.Println(reflect.TypeOf(x))

		// Insert data into session variable
		data.InsertSession(sess)

		fmt.Println(sess)

		data.UpdateStatus(user.LoggedIn, user.Username)

		// Send user information back to client using JSON format
		userInfo := data.GetUserProfile(user.Username)
		// fmt.Println(userInfo)
		js, err := json.Marshal(userInfo)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
		w.Write([]byte(js))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error: Email or password is incorrect.") // Checked in authentication.js, alerts user
	}
}

// // logout handle
func (data *Forum) LogoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LogOut Handler Here ********* ")

	c, err := r.Cookie("session_token")
	if err != nil {
		log.Fatal(err)
	}

	sess := data.GetSession(c.Value)

	fmt.Printf("User %d wants to logout\n", sess.userID)
	loggedin := "false"

	data.DeleteSession(w, sess.userID)
	data.UpdateStatus(loggedin, sess.username)

// Send user information back to client using JSON format
userInfo := data.GetUserProfile(sess.username)
// fmt.Println(userInfo)
js, err := json.Marshal(userInfo)
if err != nil {
	log.Fatal(err)
}
w.WriteHeader(http.StatusOK) // Checked in authentication.js, alerts user
w.Write([]byte(js))
}

// TODO
// once login check session table for creating user list (get data for user table)
// when logged in try to avoid loggin in another browser
