let socket;

const createWebsocket = () => {
  return new WebSocket("ws://localhost:8080/ws");
};

function showRegistrationUI() {
  document.querySelector("#registration").style.display = "flex";
  document.querySelector("#login").style.display = "none";
}

function showLoginUI() {
  document.querySelector("#login").style.display = "flex";
  document.querySelector("#registration").style.display = "none";
  
}

function showChat() {
  document.querySelector("#chat").style.display = "block";
  document.querySelector("#login").style.display = "none";
  document.querySelector("#registration").style.display = "none";
}
var conn;
var msg = document.getElementById("msg");
var log = document.getElementById("log");

function appendLog(item) {
  var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
  log.appendChild(item);
  if (doScroll) {
    log.scrollTop = log.scrollHeight - log.clientHeight;
  }
}
document.getElementById("form").onsubmit = function () {
  if (!socket) {
    return false;
  }
  if (!msg.value) {
    return false;
  }
  socket.send(msg.value);
  msg.value = "";
  return false;
};
function showFeed() {
  document.querySelector(".auth-container").style.display = "none";
  document.querySelector("main").style.display = "block";
  socket = createWebsocket();
  socket.onopen = () => {
    console.log("Socket open", socket);
  };
  socket.onmessage = function (evt) {
    var messages = evt.data.split("\n");
    console.log(messages);
    for (var i = 0; i < messages.length; i++) {
      var item = document.createElement("div");
      item.innerText = messages[i];
      appendLog(item);
    }
  };
}

// let messengerButton = document.getElementById("messenger");
// messengerButton.addEventListener("click", () => {
//   document.getElementById("log").style.display = "block";
//   document.getElementById("form").style.display = "block";
// });
