// WebSocket()
// conn = new WebSocket("ws://" + document.location.host + "/ws");
// // Listen for messages

// socket.onmessage = (event) => {

//     console.log('Message from server ', event.data);
// });
// let data = {
//   type: "message",
//   receiverId: "id1",
//   senderId: "id2",
//   content: "the message",
// };
// socket.send(JSON.stringify(data));

const createWebsocket = () => {
  return new WebSocket("ws://localhost:8080/ws");
};

let socket = createWebsocket();

// socket.send("Hello")
// socket.onmessage = (event) => {
//   console.log(event)
// }

console.log("Attempting Websocket Connection");

socket.onopen = () => {
  console.log("Successfully Connected");
  socket.send("Hi from the Client");
};

addEventListener("click", (e) => {
  // implement this on the message input
  socket.send("hello");
});
socket.onmessage = (event) => {
  console.log(event.data);
};
socket.onclose = (event) => {
  console.log("Socket Closed Connection", event);
};

socket.onerror = (error) => {
  console.error("Socket Error: ", error);
};
