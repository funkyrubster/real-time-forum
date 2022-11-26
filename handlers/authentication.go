package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

// Handles receiving the post data and adding it to the 'posts' table in the database
func (data *Forum) Post(w http.ResponseWriter, r *http.Request) {
	// Decodes posts data into post variable
	var post Post

	// Decode the JSON data from the request body into the post variable
	json.NewDecoder(r.Body).Decode(&post)

	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	// Convert data into variables for easier use
	hashtag := post.Hashtag
	time := time.Now()
	content := post.Content

	// Checks session from 'sessions' table and selects the latest one
	sess := data.GetSession()
	currentSession := sess[len(sess)-1]
	
	// Fetches username from current session
	user := currentSession.username

	type postSessionStruct struct {
		Post    []Post
		Session UserSession
	}

	// Creates postAndSession variable and assigns the post and session to it
	var postAndSession postSessionStruct
	postAndSession.Session = currentSession

	// Inserts post into the 'posts' table of the database
	data.CreatePost(Post{
		Username:  user,
		Content:   content,
		Hashtag:  hashtag,
		CreatedAt: time,
	})
	
}

// TODO: Rewrite this function to allow for hashtag count updates
func (data *Forum) Hashtag(w http.ResponseWriter, r *http.Request) {
	var hashtag Hashtag
	json.NewDecoder(r.Body).Decode(&hashtag)

	w.Write([]byte(hashtag.hashtagName))
	w.Write([]byte("ok"))
	hashName := hashtag.hashtagName
	hashCount := hashtag.hashtagCount

	sess := data.GetSession()
	currentSession := sess[len(sess)-1]

	type hashSessionStruct struct {
		Hashtag    []Hashtag
		Session UserSession
	}

	var hashtagAndSession hashSessionStruct
	hashtagAndSession.Session = currentSession

	data.GetHashtags(Hashtag{
		hashtagName:  hashName,
		hashtagCount: hashCount,
	})
	
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
			query, err1 := data.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + user.Username + "','" + user.Email + "','" + string(passwordHash) + "','" + user.Firstname + "','" + user.Lastname + "'," + user.Age + ",'" + user.Gender + "')")
			
			// Handles errors inserting data
			if err1 != nil {
				log.Fatal(err1)
			}
			_, err1 = query.Exec()
			fmt.Println(err1)

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
		fmt.Println("usID:", usID)
		fmt.Println("user.Username:", user.Username)

		// Creates a new session for the user
		sess.username = user.Username
		sess.userID = usID
		sess.max_age = 18000
		sess.session = uuid.NewV4().String()

		// Set client cookie for "session_token" as session token we just generated, also set expiry time to 120 minutes
		http.SetCookie(w, &http.Cookie{
			Name:   "session_token",
			Value:  sess.session,
			MaxAge: 900,
		})

		// Insert data into session variable
		data.InsertSession(sess)

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
// func (data *Forum) LogoutUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("LogOut Handler Here ********* ")
// 	c, err := r.Cookie("session_token")
// 	var logoutUser int

// 	if err == nil {

// 		rows, err := data.DB.Query("SELECT userID FROM sessions WHERE sessionID=?", c.Value)
// 		if err != nil {
// 			log.Fatal(err)

// 			// fmt.Println("Logout error: ", err)
// 		}
// 		defer rows.Close()
// 		for rows.Next() {
// 			rows.Scan(&logoutUser)
// 		}
// 		fmt.Printf("User %d wants to logout\n", logoutUser)
// 	}
// 	data.DeleteSession(w, logoutUser) // ?
// 	// fmt.Println("User logged out")
// 	// http.Redirect(w, r, "/", http.StatusFound)

// 	stmt, errUpdate := data.DB.Prepare("UPDATE users SET loggedin = ? WHERE userID = ?;")
// 	if errUpdate != nil {
// 		log.Fatal("Updating Table: ", errUpdate)
// 	}
// 	defer stmt.Close()
// 	stmt.Exec(false, logoutUser)
// }

// TODO
// once login check session table for creating user list (get data for user table)
// when logged in try to avoid loggin in another browser
