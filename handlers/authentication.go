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

	fmt.Print(r.Body)

	// Create user type of User struct
	var user User

// use web soc to read the information 

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("password:",user.Password)

	fmt.Println("hi from golang", user)
	w.Header().Set("Content-type", "application/text")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

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
	row = data.DB.QueryRow("select username from users where username= ?", user.Nickname)
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

	// Create user type of User struct
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-type", "application/text")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	// Only true if email/username and password match is found in the database
	emailPassCombinationValid := false
	userPassCombinationValid := false

	// Check if user entered an email or username
	enteredEmail := strings.Contains(user.Nickname, "@")
	fmt.Println(user.Nickname)
	fmt.Println(user.Password)
	
	if enteredEmail {
		fmt.Println("here")
		fmt.Println(user)
		// Check if email and password exist in users table on the same row
		var tempEmail string
		var passwordHash string
		rows, _ := data.DB.Query("SELECT email, password FROM users")

		fmt.Println("here2")
		for rows.Next(){
			rows.Scan(&tempEmail, &passwordHash)
			fmt.Println(passwordHash)
			fmt.Println(tempEmail)
     err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		 if err != nil && tempEmail == user.Nickname{
				fmt.Println("look here")
			emailPassCombinationValid = true
			fmt.Println(emailPassCombinationValid)
		}
	
	}
	} else{
		// Check if username and password exist in users table on the same row
		// if !emailPassCombinationValid {
			var tempUser string
			var passwordHash string
			rows, _ := data.DB.Query("SELECT username, password FROM users")
			for rows.Next(){
				rows.Scan(&tempUser, &passwordHash)
				err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
		if err != nil && tempUser == user.Nickname  {
			userPassCombinationValid = true
		}
	}
}
		if emailPassCombinationValid || userPassCombinationValid {
			fmt.Println("User successfully logged in.")
			// send it with web soc and
		} else {
			fmt.Println("Error: Email or password is incorrect.")
		}
}