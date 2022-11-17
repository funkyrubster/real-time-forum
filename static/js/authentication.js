<<<<<<< HEAD
let data, nickname


const signUpData = document.getElementById("sign-up-form");

signUpData.addEventListener("submit", function(){
=======
const signUpData = document.getElementById("sign-up-form");

// Create an instance of Notyf
var notyf = new Notyf();

signUpData.addEventListener("submit", function () {
>>>>>>> 6a698139129fb421214a25ddf42564246b01067d
  let user = {
    firstname: document.getElementById("firstName").value,
    lastname: document.getElementById("lastName").value,
    email: document.getElementById("email").value,
    newusername: document.getElementById("newusername").value,
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
<<<<<<< HEAD
    newpassword: document.getElementById("newpassword").value,
=======
    newpassword: document.getElementById("newpassword").value
>>>>>>> 6a698139129fb421214a25ddf42564246b01067d
  };

  let options = {
    method: "POST",
    headers: {
<<<<<<< HEAD
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch('http://localhost:8080/register', options);
  fetchRes
  .then((response) => {
    if (response.status == '200'){
     // show alert pop up  successfully created account
     alert("You have registered successfully")
      showLoginUI()
    }
    else{
      // pop up unsuccessfull
      alert("Email or username already exists")
      console.log("Email or username already exists");
    }
    return response.json();
  })
  .then((data) => {
   console.log(data);
  });
});


const loginData = document.getElementById("login-form");

loginData.addEventListener("submit", function(){
  let user = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
  }
=======
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
>>>>>>> 6a698139129fb421214a25ddf42564246b01067d

  let options = {
    method: "POST",
    headers: {
<<<<<<< HEAD
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
=======
      "Content-Type": "application/json"
    },
    body: JSON.stringify(user)
>>>>>>> 6a698139129fb421214a25ddf42564246b01067d
  };

  let fetchRes = fetch("http://localhost:8080/login", options);
  fetchRes
<<<<<<< HEAD
  .then((response) => {
    if (response.status == "200"){
      // add alert login ok
      alert("You have successfully logged in")
      showFeed()
    }
    else{
      // add alert  not ok
      alert("You inputted incorrect details")
      console.log("not ok");
    }
    return response.json()  // converted to object
  })
  .then((data) => {
    nickname = document.querySelector("#name").innerHTML = data.username
    console.log(data.username);
    // console.log(data.username);  
  })
  .catch((err) => {
    console.log(err);
  });
})

=======
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
>>>>>>> 6a698139129fb421214a25ddf42564246b01067d
