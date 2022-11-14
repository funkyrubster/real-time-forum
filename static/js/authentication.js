const signUpData = document.getElementById("sign-up-form");

signUpData.addEventListener("submit", function(){
  let user = {
    firstname: document.getElementById("firstName").value,
    lastname: document.getElementById("lastName").value,
    email: document.getElementById("email").value,
    newusername: document.getElementById("newusername").value,
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
    newpassword: document.getElementById("newpassword").value,
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch('http://localhost:8080/register', options);
  fetchRes
    .then((response) => {
      if (response.status == "200"){
        showAllCongrats()
        showLoginUI()
      }
      else{
        console.log("Email or username already exists");
      }
      // if regestration ok send to login form 
      return response.text();
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

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch("http://localhost:8080/login", options);
  fetchRes
    .then((response) => {
      console.log(response.status);
      if (response.status == '200'){
        showAlertOK()
        showHomePage()
        console.log("ok!");
      }
      else {
        console.log("its not ok");
      }
      //if login ok send to homepage 
      return response.text();
    })
    .then((data) => {
      console.log(data);
    });
});
 