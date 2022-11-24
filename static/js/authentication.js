const signUpData = document.getElementById("sign-up-form");
let allData = {};

function convertDate(date) {
  // Seperate year, day, hour and minutes
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

// Create an instance of Notyf
var notyf = new Notyf();

// ----------------- CREATE A POST -----------------
// listen for clicks on categories buttons and send the category to the server
document.querySelectorAll(".category").forEach((category) => {
  category.addEventListener("click", (e) => {
    // remove selected class from all buttons
    document.querySelectorAll(".category").forEach((category) => {
      category.classList.remove("selected");
    });
    // add selected class to the clicked button
    e.target.classList.add("selected");
    // socket.send(JSON.stringify({ category: e.target.id }));
    // console.log(category.id);
  });
});

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

function requestHashtagsUpdate() {
  console.log("requesting hashtags update");
  // Remove all posts in posts wrap
  hashtagsWrap = document.querySelector(".trending");
  hashtagsWrap.innerHTML = "";

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(user)
  };

  let fetchRes = fetch("http://localhost:8080/hashtag", options);
  fetchRes
    .then((response) => {
      return response.json();
    })
    .then(function (data) {
      allData = data;
      console.log("heres the hashtags:", data);
    })
    .catch(function (err) {
      console.log(err);
    });
}

function requestPostsUpdate() {
  // Remove all posts in posts wrap
  postsWrap = document.querySelector(".posts-wrap");
  postsWrap.innerHTML = "";

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
      return response.json();
    })
    .then(function (data) {
      allData = data;
      console.log("heres the data:", data);
      displayPosts(data);
      requestPostsUpdate(data);
    })
    .catch(function (err) {
      console.log(err);
    });
}

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
      allData = data;
      console.log("heres the data:", data);
      updateUserDetails(data);
      displayPosts(data);
      requestHashtagsUpdate();
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

// ----------------- TRENDING HASHTAGS -----------------
function displayTrendingHashtags(data) {
  console.log("displaying trending hashtags");
  console.log(data);
  // let hashtags = data.hashtags;
  // let trendingHashtags = document.querySelector(".trending-hashtags");
  // trendingHashtags.innerHTML = "";
  // hashtags.forEach((hashtag) => {
  //   trendingHashtags.innerHTML += `<a href="#" class="hashtag">#${hashtag}</a>`;
  // });
}

function displayPosts(data) {
  postsWrap = document.querySelector(".posts-wrap");

  for (let i = data.CreatedPosts.length - 1; i >= 0; i--) {
    postsWrap.innerHTML +=
      `
    <div class="post">
      <div class="header">
        <div class="author-category-wrap">
          <img src="../static/img/profile.png" width="40px" />
          <div class="name-timestamp-wrap">
            <p class="name">` +
      data.CreatedPosts[i].username +
      `</p>
            <p class="timestamp">` +
      convertDate(data.CreatedPosts[i].CreatedAt) +
      `</p>
          </div>
        </div>

        <!-- Category & Option Button -->
        <div class="category-option-wrap">
          <div class="category">` +
      data.CreatedPosts[i].Hashtag +
      `</div>
          <img src="../static/img/post-options.svg" />
        </div>
      </div>

      <!-- Post Body -->
      <div class="body">
        <p>` +
      data.CreatedPosts[i].postBody +
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

const sendPostData = function getImputValue() {
  // Get the value of the hashtag with the class of selected
  let hashtag = document.querySelector(".category.selected").innerHTML;

  let post = {
    Hashtag: hashtag,
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
    if (response.status == "200") {
      notyf.success("Your post was created successfully.");
      requestPostsUpdate();
    } else {
      notyf.error("Your post failed to send.");
    }
    return response.text();
  });
  // .then((data) =>{
  // // console.log(data);
  // });
};

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

  let fetchRes = fetch("http://localhost:8080/hashtag", options);
  fetchRes
    .then((response) => {
      return response.json();
    })
    .then(function (data) {
      allData = data;
      console.log("heres the hashtags:", data);
    })
    .catch(function (err) {
      console.log(err);
    });
});
