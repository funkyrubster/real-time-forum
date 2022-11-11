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

console.log(user);



  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch('http://localhost:8080/register', options);
  fetchRes
    .then((d) => {
      return d.text();
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
    .then((d) => {
      return d.text();
    })
    .then((data) => {
      console.log(data);
    });
});
 