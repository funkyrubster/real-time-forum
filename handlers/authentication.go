package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

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

	// Only proceed if all fields are filled
	// if len(user.Firstname) == 0 || len(user.Lastname) == 0 || len(user.Email) == 0 || len(user.Username) == 0 || len(user.Age) == 0 || len(user.Gender) == 0 || len(user.Password) == 0 {
	// 	w.WriteHeader(http.StatusNotAcceptable)
	// 	fmt.Println("Error: User did not fill all fields")
	// } else {
		// use web soc to read the information

		json.NewDecoder(r.Body).Decode(&user)

		fmt.Println("hi from golang", user)
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
				fmt.Println("bcrypt err:", err)
				return
			}
			// Insert user into database
			query, err1 := data.DB.Prepare("INSERT INTO users(username, email, password, firstname, lastname, age, gender) values('" + user.Username + "','" + user.Email + "','" + string(passwordHash) + "','" + user.Firstname + "','" + user.Lastname + "'," + user.Age + ",'" + user.Gender + "')")
			if err1 != nil {
				log.Fatal(err1)
			}
			_, err1 = query.Exec()
			fmt.Println(err1)
			fmt.Println("User successfully registered into users table.")
			w.WriteHeader(http.StatusOK)		
			
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Error: Email or username already exists.")
		}
	}

func (data *Forum) LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Create user type of LoginData struct
	var user LoginData

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)

	w.Header().Set("Content-type", "application/text")
	
	// Only true if email/username and password match is found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false
	
	// Check if user entered an email or username
	enteredEmail := strings.Contains(user.Username, "@")
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	
	if enteredEmail {
		fmt.Println(enteredEmail)
		fmt.Println("here")
		fmt.Println(user)
		// Check if email and password exist in users table on the same row
		var passwordHash string
		row := data.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Username)
		err := row.Scan(&passwordHash)
		if err != nil {
			fmt.Println("error with passwordhash")
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
			fmt.Println("error with passwordhash")
			fmt.Println(user.Username, user.Password)
		}
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err == nil {
			userPassCombinationValid = true
		}
	}
	if emailPassCombinationValid || userPassCombinationValid {
		fmt.Println("User successfully logged in.")
		// send response to js
		w.WriteHeader(http.StatusOK)
	  w.Write([]byte("ok"))			
		// set web soc
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Error: Email or password is incorrect.")
		}
	}
	