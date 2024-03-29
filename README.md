# GoChat

A scalable, real-time chat app with multimedia and AI chatbot features using Golang,
Gin, WebSocket and Redis, containerized with Docker and orchestrated via Kubernetes on AWS EKS.

Golang, Gin, MySQL, WebSocket, Redis, Docker, Kubernetes, AWS, Vue.js

![Logo](https://custom-images.strikinglycdn.com/res/hrscywv4p/image/upload/c_limit,fl_lossy,h_9000,w_1500,f_auto,q_auto/10148800/244973_960616.png)


## Features
Chating: 
Group chat: Communicate with multiple people in a single chat room.

Image/Video/Audio sending: Share multimedia files with other users.

Address-sharing: Share your location with other users in real-time.

Real-time chat: Messages are delivered instantly using WebSocket.

AI ChatBot: AI chatbot thourgh ChatGPT API V3.5

## Technologies Used
Golang: programming language used to develop the backend of the application.

Gin: web framework for Golang used to build the RESTful API of the application.

WebSocket: Protocol used to enable real-time communication between the client and server.

Redis, MySQL: Database used to store user information and messages.

Docker, Kubernetes, AWS: Tools used for deployment.

Vue.js: A JavaScript framework used to develop the frontend of the application.

## Installation
Clone the repository Copy code
```bash
  git clone https://github.com/MingyuanRen/GoChat.git
```

Install the required dependencies
```bash
  cd GoChat
  go mod download
```

Set up the MySQL database by executing the SQL script found in db/mysql.sql

Configure the application by creating a .env file based on the .env.example file.

Start the application
```bash
  go run main.go
```
## Demo
<img src="asset/demo/register.jpg" width="400" height="300" alt="Register"/>
<img src="asset/demo/login.jpg" width="400" height="300" alt="Login"/>
<img src="asset/demo/Personals.jpg" width="400" height="500" alt="Personals"/>
<img src="asset/demo/AddFriend.jpg" width="600" height="400" alt="AddFriend"/>
<img src="asset/demo/ChatPage.jpg" width="400" height="500" alt="ContactPage"/>
<img src="asset/demo/Chatting.jpg" width="600" height="400" alt="ChatPage"/>

<img src="asset/demo/extras.jpg" width="700" height="100" alt="Extra chatting functions"/>
<img src="asset/demo/functions.jpg" width="500" height="100" alt="Extra chatting functions"/>
<img src="asset/demo/creategroup.png" width="600" height="300" alt="CreateGroup"/>

## Usage
Navigate to http://localhost:8080 in your web browser.

Sign up for an account or log in if you already have one.

Create a new chat room or join an existing one.

Start chatting with other users in real-time.

