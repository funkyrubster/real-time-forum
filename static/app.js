// AUTHENTICATION
function showRegistrationUI() {
    // Shows #registration and hides #login
    document.querySelector("#registration").style.display = "flex";
    document.querySelector("#login").style.display = "none";
    document.querySelector("#mainContainer").style.display = "none";
}

function showLoginUI() {
    // Shows #login and hides #registration
    document.querySelector("#login").style.display = "flex";
    document.querySelector("#registration").style.display = "none";
    document.querySelector("#mainContainer").style.display = "none";
}

// CHAT
function showChat() {
    // Shows #chat and hides #login and #registration
    document.querySelector("#chat").style.display = "block";
    document.querySelector("#login").style.display = "none";
    document.querySelector("#registration").style.display = "none";
    document.querySelector("#mainContainer").style.display = "block";


    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onopen = () => {
        console.log('the websocket is open')
    }
    window.addEventListener('keydown', (e) =>{ 
        if (e.key === 'Enter'){
            conn.send("Welcome")
        }
    })


//     var conn;
//     var msg = document.getElementById("msg");
//     var log = document.getElementById("log");

//     function appendLog(item) {
//         var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
//         log.appendChild(item);
//         if (doScroll) {
//             log.scrollTop = log.scrollHeight - log.clientHeight;
//         }
//     }

//     document.getElementById("form").onsubmit = function () {
//         if (!conn) {
//             return false;
//         }
//         if (!msg.value) {
//             return false;
//         }
//         conn.send(msg.value);
//         msg.value = "";
//         return false;
//     };

//     if (window["WebSocket"]) {
//         conn = new WebSocket("ws://" + document.location.host + "/ws");
//         conn.onclose = function (evt) {
//             var item = document.createElement("div");
//             item.innerHTML = "<b>Connection closed.</b>";
//             appendLog(item);
//         };
//         conn.onmessage = function (evt) {
//             var messages = evt.data.split("\n");
//             for (var i = 0; i < messages.length; i++) {
//                 var item = document.createElement("div");
//                 item.innerText = messages[i];
//                 appendLog(item);
//             }
//         };
//     } else {
//         var item = document.createElement("div");
//         item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
//         appendLog(item);
//     }
 }
