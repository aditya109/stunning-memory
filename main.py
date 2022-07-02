from flask import Flask
import socket

app = Flask(__name__)

@app.route("/welcome")
def welcome():
    print("/welcome was hit successfully !")
    return "Welcome to server ! ⚡ ⛈️"

@app.route("/status")
def status():
    print("/status was hit successfully !")
    return "The server is running ! 🏃‍♂️⏩"

@app.route("/ping")
def ping():
    hostname = socket.gethostname()
    print("/ping was hit successfully !")
    return f"Hi there ! 👋 My host-id is {hostname} 🕸️"

if __name__ == "__main__":
    print(f'server running successfully at http://localhost:3000 !')
    app.run(host="0.0.0.0", port=3000, debug=True)


