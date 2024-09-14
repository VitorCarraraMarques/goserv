const HOST = "127.0.0.1"
const PORT = "8000"
let ws = new WebSocket(`ws://${HOST}:${PORT}/echo`);
ws.onopen = () => {
    console.log("opened");
};
ws.onclose = () => {
    console.log("closed");
};
ws.onerror = (err) => {
    console.log("error :", err);
};
ws.onmessage = (msg) => {
    console.log("message: ", msg);
    let display = document.getElementById("display");
    if (display === null) {
        throw "Failed to select display";
    }
    display.innerText += `\nMessage: ${msg.data}`;
};

let msgForm = document.getElementById("msgForm");
if (msgForm === null) {
    throw "Failed to read msgForm";
}
msgForm.addEventListener("submit", sendmessage);

function sendmessage(e: SubmitEvent) {
    e.preventDefault();

    let msgInput = <HTMLInputElement>document.getElementById("msgInput");
    if (msgInput === null) {
        throw "Could not load message";
    }
    ws.send(msgInput.value);
    msgInput.value = "";
}
