const signUpData = document.getElementById("sign-up-form");

// Create an instance of Notyf
var notyf = new Notyf();

function checkAgeOnlyNum(age) {
  return /^[0-9]+$/.test(age);
}

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
      // Handles missing fields
      if (response.status == "406") {
        if (user.firstname == "") {
          notyf.error("Please enter your first name.");
        } else if (user.lastname == "") {
          notyf.error("Please enter your last name.");
        } else if (user.email == "") {
          notyf.error("Please enter your email address.");
        } else if (user.newusername == "") {
          notyf.error("Please enter a username.");
        } else if (user.age == "") {
          notyf.error("Please enter your age.");
        } else if (checkAgeOnlyNum(user.age) == false) {
          notyf.error("Please enter a numerical age.");
        } else if (user.age < 18) {
          notyf.error("You must be 18 or over to register.");
        } else if (user.age > 100) {
          notyf.error("Please enter a valid age.");
        } else if (user.newpassword == "") {
          notyf.error("Please enter a password.");
        } else if (user.gender == "Gender") {
          notyf.error("Please select your gender.");
        }
        // Handles successful registration
      } else if (response.status == "200") {
        notyf.success("You have registered successfully.");
        showLoginUI();
        // Handles unsuccessful registration
      } else {
        notyf.error("The email or username already exists.");
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
      return response.json();
    })
    .then(function (data) {
      console.log(data);
      updateUserDetails(data);
      displayPosts();
    })
    .catch(function (err) {
      console.log(err);
    });

  // Concatenates the user's details within the HTML
  function updateUserDetails(data) {
    document.querySelector("p.name").innerHTML = data.User.firstName + ` ` + data.User.lastName;
    document.querySelector("p.username").innerHTML = `@` + data.User.username;
    document.querySelector("#postBody").placeholder = `What's on your mind, ` + data.User.firstName + `?`;
  }
});

function displayPosts() {
  postsWrap = document.querySelector(".posts-wrap");
  postCount = 5;

  for (let i = 0; i < postCount; i++) {
    postsWrap.innerHTML += `
    <div class="post">
      <div class="header">
        <div class="author-category-wrap">
          <img src="../static/img/profile.png" width="40px" />
          <div class="name-timestamp-wrap">
            <p class="name">FIRSTNAME LASTNAME</p>
            <p class="timestamp">TIMESTAMP</p>
          </div>
        </div>

        <!-- Category & Option Button -->
        <div class="category-option-wrap">
          <div class="category">#HASHTAG</div>
          <img src="../static/img/post-options.svg" />
        </div>
      </div>

      <!-- Post Body -->
      <div class="body">
        <p>POST BODY</p>
      </div>

      <!-- Footer -->
      <div class="footer">
        <!-- Comment, Like, Dislike -->
        <div class="actions">
          <img src="../static/img/comments-icon.svg" />
          <img src="../static/img/like-icon.svg" />
          <img src="../static/img/dislike-icon.svg" />
        </div>

        <!-- Comment, Like & Dislike Statistics -->
        <div class="stats">
          <div class="stat-wrapper">
            <img src="../static/img/post/comments-icon.svg" width="17px" />
            <p>NUM</p>
          </div>
          <div class="stat-wrapper">
            <img src="../static/img/post/likes-icon.svg" width="15px" height="13px" />
            <p>NUM</p>
          </div>
          <div class="stat-wrapper">
            <img src="../static/img/post/dislikes-icon.svg" width="17px" />
            <p>NUM</p>
          </div>
        </div>
      </div>
    </div>
    `;
  }
}

const sendPostData = function getImputValue() {
  let post = {
    postBody: document.getElementById("postBody").value
  };
  console.log(post);

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(post)
  };

  let fetchRes = fetch("http://localhost:8080/post", options);
  fetchRes.then((response) => {
    console.log(response);
    return response.text();
  });
  // .then((data) =>{
  // // console.log(data);
  // });
};
