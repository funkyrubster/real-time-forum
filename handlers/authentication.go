package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)


func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Only true if the provided email and username is not already in the database
	emailValid := false
	usernameValid := false

	var user User 


	// Get the form values
	user.Firstname = r.FormValue("first_name")
	lastname := r.FormValue("last_name")
	email := r.FormValue("email")
	username := r.FormValue("username")
	age := r.FormValue("age")
	password := r.FormValue("password")
	gender := r.FormValue("gender")

	// Open database connection
	database, _ := sql.Open("sqlite3", "database.db")

	// We need to check if there's already a user with the same username or email

	// Email check
	row := database.QueryRow("select email from users where email= ?", email)
	temp := ""
	row.Scan(&temp)
	if temp == "" {
		emailValid = true
	}

	// Username check
	row = database.QueryRow("select username from users where username= ?", user.Firstname)
	temp = ""
	row.Scan(&temp)
	if temp == "" {
		usernameValid = true
	}

	// If both email and username are valid, we can insert the user into the database
	if emailValid && usernameValid {
		// Insert user into database
		query, err1 := database.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + username + "','" + email + "','" + password + "','" + user.Firstname + "','" + lastname + "'," + age + ",'" + gender + "')")
		if err1 != nil {
			log.Fatal(err1)
		}
		_, err1 = query.Exec()
		fmt.Println(err1)
	
		fmt.Println("User successfully registered into users table.")
		} else {
			fmt.Println("Error: Email or username already exists.")
		}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("enything")
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Only true if email/username and password match is found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	// Get the form values
	email := r.FormValue("emailusername")
	username := r.FormValue("emailusername")
	password := r.FormValue("password")

	// Open database connection
	database, _ := sql.Open("sqlite3", "database.db")

	// Check if user entered an email or username
	enteredEmail := strings.Contains(email, "@")

	if enteredEmail {
		// Check if email and password exist in users table on the same row
		rows, _ := database.Query("SELECT email, password FROM users")
		var tempEmail string
		var tempPassword string

		for rows.Next() {
			rows.Scan(&tempEmail, &tempPassword)
			if tempEmail == email && tempPassword == password {
				emailPassCombinationValid = true
			}
		}
	} else {
		// Check if username and password exist in users table on the same row
		if !emailPassCombinationValid {
			rows, _ := database.Query("SELECT username, password FROM users")
			var tempUsername string
			var tempPassword string

			for rows.Next() {
				rows.Scan(&tempUsername, &tempPassword)
				if tempUsername == username && tempPassword == password {
					userPassCombinationValid = true
				}
			}
		}
	}

	if emailPassCombinationValid || userPassCombinationValid {
		fmt.Println("User successfully logged in.")
	} else {
		fmt.Println("Error: Email or password is incorrect.")
	}
}
