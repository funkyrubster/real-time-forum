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

function showFeed() {
  document.querySelector(".auth-container").style.display = "none";
  document.querySelector("main").style.display = "block";
  socket = createWebsocket();
  console.log("socket mainpage", socket);
}

// listen for clicks on categories buttons and send the category to the server
document.querySelectorAll(".category").forEach((category) => {
  category.addEventListener("click", (e) => {
    // remove selected class from all buttons
    document.querySelectorAll(".category").forEach((category) => {
      category.classList.remove("selected");
    });
    // add selected class to the clicked button
    e.target.classList.add("selected");

    socket.send(JSON.stringify({ category: e.target.id }));
    console.log(category.id);
  });
});
