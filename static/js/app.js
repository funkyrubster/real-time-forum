// AUTHENTICATION
function showRegistrationUI() {
    // Shows #registration and hides #login
    document.querySelector("#registration").style.display = "flex";
    document.querySelector("#login").style.display = "none";
}

function showLoginUI() {
    // Shows #login and hides #registration
    document.querySelector("#login").style.display = "flex";
    document.querySelector("#registration").style.display = "none";
}

// CHAT
function showChat() {
    // Shows #chat and hides #login and #registration
    document.querySelector("#chat").style.display = "block";
    document.querySelector("#login").style.display = "none";
    document.querySelector("#registration").style.display = "none";
}


function showHomePage() {
    document.querySelector(".auth-container").style.display = "none";
}
