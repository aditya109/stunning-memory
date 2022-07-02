const express = require("express");
// Allocating os module
const os = require("os");
const app = express();
const port = 3000;

app.get("/welcome", (req, res) => {
  console.log("/welcome was hit successfully !");
  res.send("Welcome to server ! ⚡ ⛈️");
});

app.get("/status", (req, res) => {
  console.log("/status was hit successfully !");
  res.send("The server is running ! 🏃‍♂️⏩");
});

app.get("/ping", (req, res) => {
  console.log("/ping was hit successfully !");
  if (os.hostname()) {
    var hostname = os.hostname();
    res.send(`Hi there ! 👋 My host-id is ${hostname} 🕸️`)
  } else {
    throw new Error('hostname not found !')
  }
});

app.listen(port, () => {
  console.log(`server running successfully at http://localhost:${port} !`);
});
