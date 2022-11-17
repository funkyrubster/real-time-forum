const signUpData = document.getElementById("sign-up-form");

// Create an instance of Notyf
var notyf = new Notyf();

signUpData.addEventListener("submit", function () {
  let user = {
    firstname: document.getElementById("firstName").value,
    lastname: document.getElementById("lastName").value,
    email: document.getElementById("email").value,
    newusername: document.getElementById("newusername").value,
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
    newpassword: document.getElementById("newpassword").value
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(user)
  };

  let fetchRes = fetch("http://localhost:8080/register", options);
  fetchRes
    .then((response) => {
      if (response.status == "406") {
        // show alert pop up  successfully created account
        notyf.error("Please fill in all required fields.");
      } else if (response.status == "200") {
        // show alert pop up  successfully created account
        notyf.success("You have successfully registered.");
        showLoginUI();
      } else {
        // pop up unsuccessfull
        notyf.error("The email or username already exists.");
        console.log("Email or username already exists");
      }
      return response.text();
    })
    .then((data) => {
      console.log(data);
    });
});

const loginData = document.getElementById("login-form");

loginData.addEventListener("submit", function () {
  let user = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(user)
  };

  let fetchRes = fetch("http://localhost:8080/login", options);
  fetchRes
    .then((response) => {
      if (response.status == "200") {
        // add alert login ok
        notyf.success("You have logged in successfully.");
        // alert("You have successfully logged in");
        showFeed();
      } else {
        // add alert  not ok
        notyf.error("The login details you entered are incorrect.");
        console.log("not ok");
      }
      return response.text();
    })
    .then((data) => {
      console.log(data);
    });
});
