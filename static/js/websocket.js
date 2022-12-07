// export const createWebsocket = () => {
//   return new WebSocket("ws://localhost:8080/ws");
// };

// let socket = createWebsocket();

// console.log("Attempting websocket connection...");

// socket.onopen = () => {
//   console.log("Successfully connected.");
// };

// socket.onmessage = (event) => {
//   console.log(event.data);
// };
// socket.onclose = (event) => {
//   console.log("Socket closed connection:", event);
// };

// socket.onerror = (error) => {
//   console.error("Socket error:", error);
// };

// var conn;
// var msg = document.getElementById("msg");
// var log = document.getElementById("log");

// function appendLog(item) {
//   var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
//   log.appendChild(item);
//   if (doScroll) {
//     log.scrollTop = log.scrollHeight - log.clientHeight;
//   }
// }

// document.getElementById("form").onsubmit = function () {
//   if (!conn) {
//     return false;
//   }
//   if (!msg.value) {
//     return false;
//   }
//   conn.send(msg.value);
//   msg.value = "";
//   return false;
// };

// if (window["WebSocket"]) {
//   conn = new WebSocket("ws://" + document.location.host + "/ws");
//   conn.onclose = function (evt) {
//     var item = document.createElement("div");
//     item.innerHTML = "<b>Connection closed.</b>";
//     appendLog(item);
//   };
//   conn.onmessage = function (evt) {
//     var messages = evt.data.split('\n');
//     for (var i = 0; i < messages.length; i++) {
//       var item = document.createElement("div");
//       item.innerText = messages[i];
//       appendLog(item);
//     }
//   };
// } else {
//   var item = document.createElement("div");
//   item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
//   appendLog(item);
// }

// let messengerButton = document.getElementById("messenger");
// messengerButton.addEventListener("click", () => {
//   document.getElementById("log").style.display = "block";
//   document.getElementById("form").style.display = "block";
// })