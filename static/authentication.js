const signUpData = document.getElementById("sign-up-form");

signUpData.addEventListener("submit", function(){
  let user = {
    username: document.getElementById("username").value,
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
    firstname: document.getElementById("firstName").value,
    lastname: document.getElementById("lastName").value,
    email: document.getElementById("email").value,
    password: document.getElementById("password").value,
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch("/register", options);
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
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
    firstname: document.getElementById("firstName").value,
    lastname: document.getElementById("lastName").value,
    email: document.getElementById("email").value,
    password: document.getElementById("password").value,
  }

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  };

  let fetchRes = fetch("/login", options);
  fetchRes
    .then((d) => {
      return d.text();
    })
    .then((data) => {
      console.log(data);
    });
});
 