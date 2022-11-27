// Used for sending notifications
var notyf = new Notyf();

// Used for converting the date to a more readable format
function convertDate(date) {
  // Seperate year, day, hour and minutes into vars
  let yyyy = date.slice(0, 4);
  let dd = date.slice(8, 10);
  let hh = date.slice(11, 13);
  let mm = date.slice(14, 16);

  // Get int for day of the week (0-6, Sunday-Saturday)
  const d = new Date(date);
  let dayInt = d.getDay();
  let day = "";
  switch (dayInt) {
    case 0:
      day = "Sunday";
      break;
    case 1:
      day = "Monday";
      break;
    case 2:
      day = "Tuesday";
      break;
    case 3:
      day = "Wednesday";
      break;
    case 4:
      day = "Thursday";
      break;
    case 5:
      day = "Friday";
      break;
    case 6:
      day = "Saturday";
      break;
  }

  // Get int for month (0-11, January-December)
  let monthInt = d.getMonth();
  let month = "";
  switch (monthInt) {
    case 0:
      month = "January";
      break;
    case 1:
      month = "February";
      break;
    case 2:
      month = "March";
      break;
    case 3:
      month = "April";
      break;
    case 4:
      month = "May";
      break;
    case 5:
      month = "June";
      break;
    case 6:
      month = "July";
      break;
    case 7:
      month = "August";
      break;
    case 8:
      month = "September";
      break;
    case 9:
      month = "October";
      break;
    case 10:
      month = "November";
      break;
    case 11:
      month = "December";
      break;
  }
  fullDate = day + ", " + dd + " " + month + ", " + yyyy + " @ " + hh + ":" + mm;
  return fullDate;
}

/* ---------------------------------------------------------------- */
/*                         REGISTERING USERS                        */
/* ---------------------------------------------------------------- */
const signUpData = document.getElementById("sign-up-form");
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
  fetchRes.then((response) => {
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
  });
});

// Used for validating age field on sign up
function checkAgeOnlyNum(age) {
  return /^[0-9]+$/.test(age);
}

/* ---------------------------------------------------------------- */
/*                       AUTHENTICATING USERS                       */
/* ---------------------------------------------------------------- */
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
      }
      return response.json();
    })
    .then(function (data) {
      // Fills the user's profile with their details
      updateUserDetails(data);
      // Pulls latest posts from database and displays them
      refreshPosts();
      // Pulls hashtag stats from database and displays them
      refreshHashtags();
    })
    .catch(function (err) {
      console.log(err);
    });
});

// Concatenates the user's details within the HTML after login
function updateUserDetails(data) {
  document.querySelector("p.name").innerHTML = data.User.firstName + ` ` + data.User.lastName;
  document.querySelector("p.username").innerHTML = `@` + data.User.username;
  document.querySelector("#postBody").placeholder = `What's on your mind, ` + data.User.firstName + `?`;
}

function refreshPosts() {
  fetch("/getPosts", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    },
    method: "POST"
  })
    .then((response) => {
      response.text().then(function (data) {
        let posts = JSON.parse(data);
        console.log("posts:", posts);
        // 'posts' contains all latest posts from database, in JSON format
        displayPosts(posts);
      });
    })
    .catch((error) => {
      console.log(error);
    });
}

function refreshHashtags() {
  fetch("/getHashtags", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    },
    method: "POST"
  })
    .then((response) => {
      response.text().then(function (data) {
        let hashtags = JSON.parse(data);
        console.log("Hashtags data:", hashtags);
        // 'hashtags' contains all latest hashtags & counts from database, in JSON format
        displayTrendingHashtags(hashtags);
      });
    })
    .catch((error) => {
      console.log(error);
    });
}

// Listen for clicks on categories buttons and adds 'selected' class
document.querySelectorAll(".category").forEach((category) => {
  category.addEventListener("click", (e) => {
    // remove selected class from all buttons
    document.querySelectorAll(".category").forEach((category) => {
      category.classList.remove("selected");
    });
    // add selected class to the clicked button
    e.target.classList.add("selected");
  });
});

// Sends the user's post to the server
const createPost = function getInputValue() {
  // Get the value of the hashtag with the class of selected
  let hashtag = document.querySelector(".category.selected").innerHTML;

  let post = {
    postBody: document.getElementById("postBody").value,
    Hashtag: hashtag
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(post)
  };

  let fetchRes = fetch("http://localhost:8080/post", options);
  fetchRes.then((response) => {
    if (response.status == "200") {
      notyf.success("Your post was created successfully.");
      refreshPosts();
    } else {
      notyf.error("Your post failed to send.");
    }
    return response.text();
  });
};

// Displays all posts on the feed
function displayPosts(posts) {
  postsWrap = document.querySelector(".posts-wrap");

  // Clear all posts printed
  postsWrap.innerHTML = "";

  // Loop through all posts and print them, concatenating each post data
  for (let i = posts.length - 1; i >= 0; i--) {
    postsWrap.innerHTML +=
      `
    <div class="post">
      <div class="header">
        <div class="author-category-wrap">
          <img src="../static/img/profile.png" width="40px" />
          <div class="name-timestamp-wrap">
            <p class="name">` +
      posts[i].username +
      `</p>
            <p class="timestamp">` +
      convertDate(posts[i].CreatedAt) +
      `</p>
          </div>
        </div>
        <!-- Category & Option Button -->
        <div class="category-option-wrap">
          <div class="category">` +
      posts[i].Hashtag +
      `</div>
          <img src="../static/img/post-options.svg" />
        </div>
      </div>
      <!-- Post Body -->
      <div class="body">
        <p>` +
      posts[i].postBody +
      `</p>
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
            <p>0</p>
          </div>
          <div class="stat-wrapper">
            <img src="../static/img/post/likes-icon.svg" width="15px" height="13px" />
            <p>0</p>
          </div>
          <div class="stat-wrapper">
            <img src="../static/img/post/dislikes-icon.svg" width="17px" />
            <p>0</p>
          </div>
        </div>
      </div>
    </div>
    `;
  }
}

// Displays all posts on the feed
function displayTrendingHashtags(hashtags) {
  trendingWrap = document.querySelector(".trending");

  // Clear all hashtags printed
  trendingWrap.innerHTML = "";

  // Loop through all hashtags and print them, concatenating each hashtag data
  for (let i = hashtags.length - 1; i >= 0; i--) {
    trendingWrap.innerHTML +=
      `
      <div class="hashtag">
        <p id="name">` +
      hashtags[i].name +
      `</p>
        <div class="circle">
          <p id="count">` +
      hashtags[i].count +
      `</p>
        </div>
      </div>
    `;
  }
}
