package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

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

func (data *Forum) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Create user type of RegisterData struct
	var user RegisterData
	json.NewDecoder(r.Body).Decode(&user)

	numChars := "0123456789"

	// Only proceed if all fields are filled
	if len(user.Firstname) == 0 || len(user.Lastname) == 0 || len(user.Email) == 0 || len(user.Username) == 0 || (len(user.Age) == 0 || !strings.ContainsAny(user.Age, numChars)) || user.Gender == "Gender" || len(user.Password) == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		// use web soc to read the information
		w.Header().Set("Content-type", "application/text")

		// // Only true if the provided email and username is not already in the database
		emailValid := false
		usernameValid := false

		// We need to check if there's already a user with the same username or email

		// Email check
		row := data.DB.QueryRow("select email from users where email= ?", user.Email)
		temp := ""
		row.Scan(&temp)
		if temp == "" {
			emailValid = true
		}

		// Username check
		row = data.DB.QueryRow("select username from users where username= ?", user.Username)
		temp = ""
		row.Scan(&temp)
		if temp == "" {
			usernameValid = true
		}

		// If both email and username are valid, we can insert the user into the database
		if emailValid && usernameValid {

			var passwordHash []byte

			// create hash from password
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Error hashing password:", err)
				return
			}
			// Insert user into database
			query, err1 := data.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + user.Username + "','" + user.Email + "','" + string(passwordHash) + "','" + user.Firstname + "','" + user.Lastname + "'," + user.Age + ",'" + user.Gender + "')")
			if err1 != nil {
				log.Fatal(err1)
			}
			_, err1 = query.Exec()
			fmt.Println(err1)

			fmt.Println("SUCCESS: User successfully registered into users table.")
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("ERROR: Username or email already exists.")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (data *Forum) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var sess UserSession

	// Create user type of LoginData struct
	var user LoginData
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-type", "application/text")

	// Only true if email/username and password match is found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	// Check if user entered an email or username
	enteredEmail := strings.Contains(user.Username, "@")

	if enteredEmail {
		// Check if email and password exist in users table on the same row
		var passwordHash string
		row := data.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Username)
		err := row.Scan(&passwordHash)
		if err != nil {
			fmt.Println("Error with password hash:", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err == nil {
			emailPassCombinationValid = true
		}
	} else {
		// Check if username and password exist in users table on the same row
		var passwordHash string
		row := data.DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username)
		err := row.Scan(&passwordHash)
		if err != nil {
			fmt.Println("Error with password hash:", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err == nil {
			userPassCombinationValid = true
		}
	}
	var usID int = 5
	if emailPassCombinationValid || userPassCombinationValid {
		fmt.Println("User logged in successfully.")

		row := data.DB.QueryRow("SELECT userID FROM users WHERE username = ?;", user.Username)
		err := row.Scan(&usID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("usID:", usID)
		fmt.Println("user.Username:", user.Username)
		sess.userID = usID
		sess.max_age = 18000
		sess.session = uuid.NewV4().String()

		// Finally, we set the client cookie for "session_token" as the session token we just generated
		// we also set an expiry time of 120 minutes
		http.SetCookie(w, &http.Cookie{
			Name:   "session_token",
			Value:  sess.session,
			MaxAge: 900,
		})
		// insert into session
		data.InsertSession(sess)

		// send response to js
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		// set web soc
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error: Email or password is incorrect.")
	}
}

// InsertSession ...
func (data *Forum) InsertSession(sess UserSession) {
	stmnt, err := data.DB.Prepare("INSERT INTO sessions (cookieValue, userID) VALUES (?, ?)")
	if err != nil {
		fmt.Println("AddSession error inserting into DB: ", err)
	}
	defer stmnt.Close()
	stmnt.Exec(sess.session, sess.userID)
}

// User's cookie expires when browser is closed, delete the cookie from the database.
func (data *Forum) DeleteSession(w http.ResponseWriter, userID int) error {
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	stmt, err := data.DB.Prepare("DELETE FROM session WHERE userID=?;")
	defer stmt.Close()
	stmt.Exec(userID)
	if err != nil {
		fmt.Println("DeleteSession err: ", err)
		return err
	}
	return nil
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
