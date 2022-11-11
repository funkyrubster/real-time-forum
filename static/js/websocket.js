let socket = new WebSocket("ws://localhost:8080/ws")

console.log("Attempting Websocket Connection")

socket.onopen = () => {
  console.log("Successfully Connected");
  socket.send("Hi from the Client")
}


socket.onclose = (event) => {
  console.log("Socket Closed Connection",event);
}

socket.onerror = (error) => {
  console.error("Socket Error: ", error)
  
}; 