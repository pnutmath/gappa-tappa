const socket = new WebSocket("ws://localhost:8080/ws")

let connect = cb => {
    console.log("Attempting connection...")

    socket.onopen = () => {
        console.log("Successfully Connected")
    }

    socket.onmessage = msg => {
        cb(msg);
    }

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    }

    socket.onerror = error => {
        console.error("Socket error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
}

export { connect, sendMsg }