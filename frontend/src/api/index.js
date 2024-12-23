var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (cb) => {
  console.log("Attempting connection...");

  socket.onopen = () => {
    console.log("Connected to server");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log("Connection closed: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg) => {
  console.log("Sending message: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
