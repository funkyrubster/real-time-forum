package handlers

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func (data *Forum) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Only true if the provided email and username is not already in the database
	emailValid := false
	usernameValid := false

	// Create user type of User struct
	var user User

	// Get the form values
	user.Firstname = r.FormValue("first_name")
	user.Lastname = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Nickname = r.FormValue("username")
	user.Age = r.FormValue("age")
	user.Password = r.FormValue("password")
	user.Gender = r.FormValue("gender")

	// We need to check if there's already a user with the same username or email

	// Email check
	row := data.DB.QueryRow("select email from users where email= ?", user.Email)
	temp := ""
	row.Scan(&temp)
	if temp == "" {
		emailValid = true
	}

	// Username check
	row = data.DB.QueryRow("select username from users where username= ?", user.Firstname)
	temp = ""
	row.Scan(&temp)
	if temp == "" {
		usernameValid = true
	}

	// create hash from password

	var passwordHash []byte

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return
	}

	// If both email and username are valid, we can insert the user into the database
	if emailValid && usernameValid {
		// Insert user into database
		query, err1 := data.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + user.Nickname + "','" + user.Email + "','" + string(passwordHash) + "','" + user.Firstname + "','" + user.Lastname + "'," + user.Age + ",'" + user.Gender + "')")
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

func (data *Forum) LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Create user type of User struct
	var user User

	// Only true if email/username and password match is found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	var passwordHash string
	var tempEmail string
	var tempUser string

	// Get the form values
	user.Email = r.FormValue("emailusername")
	user.Nickname = r.FormValue("emailusername")
	user.Password = r.FormValue("password")

	// Check if email and password exist in users table on the same row
	rows := data.DB.QueryRow("SELECT email, password FROM users")
	err := rows.Scan(&tempEmail, &passwordHash)
	fmt.Println("hash from db:", passwordHash)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username")
		return
	}
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))

	if err == nil {
		emailPassCombinationValid = true
	} else {

		// Check if username and password exist in users table on the same row
		if !emailPassCombinationValid {
			rows := data.DB.QueryRow("SELECT username, password FROM users")
			err := rows.Scan(&tempUser, &passwordHash)
			if err != nil {
				log.Fatal(err)
				return
			}
			if err == nil {
				userPassCombinationValid = true
			}
		}
		if emailPassCombinationValid || userPassCombinationValid {
			fmt.Println("User successfully logged in.")
		} else {
			fmt.Println("Error: Email or password is incorrect.")
		}
	}
}
