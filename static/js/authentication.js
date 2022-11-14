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
    return response.text();
  })
  .then((data) => {
    
    console.log(data);
  });
});

