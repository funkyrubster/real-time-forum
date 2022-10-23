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
