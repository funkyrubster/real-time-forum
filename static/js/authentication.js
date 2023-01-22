// Used for sending notifications
var notyf = new Notyf();

function convertTime(date) {
  // Seperate year, day, hour and minutes into vars
  let yyyy = date.slice(0, 4);
  let dd = date.slice(8, 10);
  let hh = date.slice(11, 13);
  let mm = date.slice(14, 16);

  output = hh + ":" + mm;
  return output;
}

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
let userData;
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
      userData = data;
      onlineActivity();
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
  // console.log(
  //   "updated User Details",
  //   data.User.firstName,
  //   data.User.lastName,
  //   data.User.username
  // );
  document.querySelector("p.name").innerHTML = data.User.firstName + ` ` + data.User.lastName;
  document.querySelector("p.username").innerHTML = `@` + data.User.username;
  document.querySelector("p.username").setAttribute("data-userId", data.User.userID);
  document.querySelector("#postBody").placeholder = `What's new, ` + data.User.firstName + `?`;
}

function onlineActivity() {
  fetch("/usersStatus", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    },
    method: "POST"
  })
    .then((response) => {
      response.text().then(function (data) {
        let status = JSON.parse(data);

        userActivityWrapper = document.querySelector("#recently-joined > div");

        userActivityWrapper.innerHTML = "";
        // console.log(status);
        for (let i = 0; i < status.Online.length; i++) {
          userActivityWrapper.innerHTML +=
            `
            <div class="user" data-reciverid="${status.Online[i].userID}" onclick="startChat(${i}, ${status.Online[i].userID})">
              <div class="online-status"></div>
              <p>` +
            status.Online[i].username +
            `</p>
            <div class="notification" id = "${status.Online[i].username + "-notification"}">notification here
            </div>
            </div>
      `;
        }
        let name = document.querySelector(".name").textContent;
        let username = Array.from(document.querySelectorAll(".fullname"));
        for (let i = 0; i < username.length; i++) {
          if (name == username[i].textContent) {
            username[i].style.display = "none";
          }
        }
        if (status.Offline == null) {
          // console.log("empty");
        } else {
          for (let i = 0; i < status.Offline.length; i++) {
            // console.log(status.Offline[i].firstName);
            userActivityWrapper.innerHTML +=
              `
            <div class="user" data-reciverid="${status.Online[i].userID}" onclick="startChat(${i}, ${status.Online[i].userID})">
              <div class="offline-status"></div>
              <p>` +
              status.Offline[i].firstName +
              " " +
              status.Offline[i].lastName +
              `</p>
            </div>
            `;
          }
        }
      });
    })
    .catch((error) => {
      console.log(error);
    });
}

function toggleChat() {
  chatDiv = document.querySelector(".chat");

  if (chatDiv.style.display === "flex") {
    chatDiv.style.display = "none";
    chatDiv.classList.remove("show");
  } else {
    chatDiv.style.display = "flex";
    chatDiv.classList.add("show");
  }
}

// ISSUE: if we click an offline user, is recieves userid 1 private message.

function startChat(index, id) {
  onlineActivity();
  let arrayOfOnlineUsers = Array.from(document.querySelectorAll(".user p")); // changed here
  let fullName = arrayOfOnlineUsers[index].textContent;
  document.querySelector("#chat > div.profile-header > div > p").innerHTML = fullName;
  document.querySelector("#chat > div.profile-header > div > p").setAttribute("data-reciverid", id);
  //document.querySelector("#online > li").dataset.reciverid
  let sendername = document.querySelector("#username-id").textContent;
  console.log(sendername);
  let newStr = sendername.replace("@", "");
  let senderuser = {
    sendersusername: newStr,
    recipientsusername: fullName
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(senderuser)
  };

  let fetchRes = fetch("http://localhost:8080/loadingmessage", options);

  fetchRes
    .then((response) => {
      return response.json();
    })
    // console.log("starting fetch")
    .then(function (data) {
      // problem solved. Code wasn't reachable beause of print statement above.

      // let messages = JSON.parse(data);
      console.log("Messages here: ", data);
      document.querySelector("#log").innerHTML = ""; // clears the chat box
      displayMessages(data);
    });
}

// check the current position and then the next after the next 10 is loaded.
//

function displayMessages(messages) {
  console.log("INSIDE display: ", messages);
  let start = document.querySelector("#log").childElementCount;

  let currUser = document.querySelector("#username-id").textContent;
  console.log(currUser);

  document.querySelector("#log").innerHTML = "";

  messages = messages.reverse(); // gets the latest 10 messages in database
  for (let i = start; i < start + 10; i++) {
    if (currUser.slice(1) !== messages[i].messagesender) {
      document.querySelector("#log").innerHTML += `
        <div class="bubbleWrapper">
          <div class="inlineContainer">
            <div class="otherBubble other">
              ${messages[i].message}
            </div>
          </div>
          <span class="other">
            ${convertTime(messages[i].CreatedAt)}
          </span>
        </div>
    `;
    } else {
      console.log(messages[i].messagesender);
      document.querySelector("#log").innerHTML += `
      <div class="bubbleWrapper">
        <div class="inlineContainer own">
          <div class="ownBubble own">
            ${messages[i].message} 
          </div>
        </div>
        <span class="own">
          ${convertTime(messages[i].CreatedAt)}
        </span>
      </div>
      `;
    }
  }
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
        // console.log("posts:", posts);
        // 'posts' contains all latest posts from database, in JSON format
        displayPosts(posts);
      });
    })
    .catch((error) => {
      console.log(error);
    });
}

function refreshComments(postID) {
  let commentData = {
    postId: postID
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(commentData)
  };
  let fetchRes = fetch("http://localhost:8080/sendComments", options);
  fetchRes
    .then((response) => {
      return response.json();
    })
    .then(function (data) {
      // sends latest comment data to getComments function
      getComments(data, postID);
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

const saveChat = function getChatContents() {
  // console.log(document.getElementById("log"));

  let receiver = document.querySelector("#chatReceiver").textContent;
  let chat = {
    message: document.getElementById("msg").value,
    messagerecipient: receiver
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(chat)
  };

  let fetchResChat = fetch("http://localhost:8080/chat", options);
  fetchResChat.then((response) => {
    return response.text();
  });
  // showChat();
};

// Sends the user's post to the server
function createPost() {
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
      postBody.value = "";
      notyf.success("Your post was created successfully.");
      refreshPosts();
      updateHashtagTable();
    } else {
      notyf.error("Your post failed to send.");
    }
    return response.text();
  });
}

// Displays all posts on the feed
function displayPosts(posts) {
  postsWrap = document.querySelector(".posts-wrap");

  // Clear all posts printed
  postsWrap.innerHTML = "";

  // Loop through all posts and print them, concatenating each post data
  for (let i = (posts ? posts.length : 0) - 1; i >= 0; i--) {
    postsWrap.innerHTML +=
      `
    <div class="post" id="` +
      posts[i].PostID +
      `">
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
        <!-- Category Button -->
        <div class="category">` +
      posts[i].Hashtag +
      `</div>
      </div>
      <!-- Post Body -->
      <div class="body">
        <p>` +
      posts[i].postBody +
      `</p>
      </div>
      <!-- Footer -->
      <!-- Footer -->
      <div class="footer">
        <!-- Comment, Like, Dislike -->
        <div class="actions">
          <img src="../static/img/comments-icon.svg" onclick="refreshComments(${posts[i].PostID})" id="${posts[i].PostID}"/>
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

      <div class="comments">
                <!-- Create A Comment -->
                <div class="separator"></div>
                <div class="create-comment-wrap">
                  <div class="comment-field-wrap">
                    <!-- <img src="../static/img/profile.png" width="50px" id="composeCommentAuthor"> -->
                    <div class="comment-field-submit-wrap">
                      <input type="text" id="commentBody${posts[i].PostID}" placeholder="Write a comment..." />
                      <div class="comment-btn" onclick="createCom(${posts[i].PostID})">Comment</div>
                    </div>
                  </div>
                </div>
                <div class="separator"></div>

                <!-- Comments -->
                <p class="title">Comments</p>
                <div class="comments-wrap">
                  <img src="../static/img/post/comments/no-comments.svg" width="600px" />
                </div>
              </div>
    </div>
    `;
  }
}

function createCom(postID) {
  console.log("postID: " + postID);
  let idCommentBody = "#commentBody" + postID;
  let comBody = document.querySelector(idCommentBody);

  let commentObj = {
    postid: postID,
    commentBody: comBody.value
  };
  console.log(commentObj);
  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(commentObj)
  };
  let fetchRes = fetch("http://localhost:8080/comment", options);
  fetchRes.then((response) => {
    if (response.status == "200") {
      notyf.success("Your comment was created successfully.");
      refreshComments(postID);
      comBody.value = " ";
    } else {
      notyf.error("Your comment failed to send.");
    }
    return response.text();
  });
}

function getComments(comments, postID) {
  // update comments counter
  let commentsCounter = document.querySelector("#\\3" + postID + "  > div.footer > div.stats > div:nth-child(1) > p");
  commentsCounter.innerHTML = comments.length;

  console.log(comments);
  console.log("first com", comments[1]);

  commentsWrap = document.querySelector("#\\3" + postID + "  > div.comments > div.comments-wrap");

  // Clear all posts printed
  commentsWrap.innerHTML = "";

  // Loop through all comments and print them
  for (let i = comments.length - 1; i >= 0; i--) {
    commentsWrap.innerHTML +=
      `
        <div class="comment">
        <div class="author">` +
      comments[i].username +
      `</div>
        <img src="../static/img/profile.png" id="profile-picture" width="35px">
        <div class="timestamp">` +
      convertDate(comments[i].CreatedAt) +
      `</div>
        <div class="body">` +
      comments[i].commentBody +
      `</div>
      </div>
          `;
  }
}

function updateHashtagTable() {
  // Get the value of the hashtag with the class of selected
  let hashtag_value = document.querySelector(".category.selected").innerHTML;

  let hashtag = {
    Name: hashtag_value,
    Count: "1"
  };

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(hashtag)
  };

  let fetchRes = fetch("http://localhost:8080/updateHashtag", options);
  fetchRes.then((response) => {
    if (response.status == "200") {
      refreshHashtags();
    } else {
      notyf.error("Failed to update trending hashtags.");
    }
    return response.text();
  });
}

// Displays all posts on the feed
function displayTrendingHashtags(hashtags) {
  // console.log(hashtags);
  trendingWrap = document.querySelector(".trending");

  // We need to check if there are any hashtag stats to print, otherwise leave at default order

  // Assume all hashtags have 0 count
  let allZero = true;

  // Check if all hashtag counts are 0 aka no posts have been made
  for (let i = 0; i <= hashtags.length - 1; i++) {
    if (hashtags[i].count != 0) {
      allZero = false;
    }
  }

  // If any hashtag count is > 0, rearrange trending div as we know there are posts
  if (!allZero) {
    // Clear existing hashtags div contents
    trendingWrap.innerHTML = "";

    // Sort hashtags by count
    hashtags.sort((a, b) => (a.count < b.count ? 1 : -1));

    // Loop through all hashtags and print them, concatenating each hashtag data
    for (let i = 0; i <= hashtags.length - 1; i++) {
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
}

const logout = function logoutUser() {
  let cookie = document.cookie;
  let username = cookie.split("=")[0];

  console.log(username);

  let logoutData = {
    ok: ""
  };

  logoutData.ok = username;

  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(logoutData)
  };

  let fetchRes = fetch("http://localhost:8080/logout", options);
  fetchRes
    .then((response) => {
      if (response.status === 200) {
        console.log("ok");
      }
      return response.json();
    })
    .then(function (data) {
      if (data.User.LoggedIn === "false") {
        document.querySelector("main").style.display = "none";
        document.querySelector(".auth-container").style.display = "flex";

        // showRegistrationUI()
        notyf.success("Succesfully logged out.");
      }
    })
    .catch(function (err) {
      console.log(err);
    });
  socket.close();
};

// Get a reference to the posts wrapper div
const postsWrapper = document.querySelector(".posts-wrap");

// Listen for clicks on the posts wrapper div
postsWrapper.addEventListener("click", (event) => {
  console.log(event.target);
  // Check if the clicked element is a post, header, body, or footer
  if (event.target.matches("img, .name, .timestamp, .category-option-wrap, .post, .body, .stat-wrapper, .stats, .author, p, .create-comment-wrap, .header, .footer")) {
    // Save the ID of the clicked post to a variable
    const clickedPostId = event.target.id;

    // Get a reference to the .comments child inside of the clicked post
    const comments = event.target.closest(".post").querySelector(".comments");

    // Check if the comments element exists
    if (comments) {
      // Check if the comments element is already visible
      if (comments.style.display === "block") {
        // If the comments element is already visible, set its display property to 'none'
        comments.style.display = "none";
      } else {
        // If the comments element is not visible, set its display property to 'block'
        comments.style.display = "block";
      }
    }
  }
});

function checkCookies() {
  let cookie = document.cookie;

  let cookieValue = document.cookie.split("=")[1];

  if (cookie != "") {
    let data = {
      cookieValue: cookieValue
    };

    let options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    };

    fetch("http://localhost:8080/checkCookie", options)
      .then((response) => response.json())
      .then((data) => {
        updateUserDetails(data);
        showFeed();
        onlineActivity();
        refreshPosts();
        refreshHashtags();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  } else {
    showLoginUI();
  }
}
