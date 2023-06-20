README - For'Rhum Arrangé
Introduction
For'Rhum Arrangé is a web forum application designed for discussions and interactions. This README file provides instructions on how to set up and run the forum application on your local machine.

Prerequisites
Before you begin, ensure that you have the following installed:

Go programming language (https://golang.org/dl/)
WSL (Windows Subsystem for Linux) if you're using Windows (https://docs.microsoft.com/en-us/windows/wsl/install-win10)
Getting Started
Clone the repository to your local machine:
Copy code
$ git clone https://github.com/Razner/forum.git
Navigate to the server directory:
Copy code
$ cd Front/Server
Launch the server using the following command:
Copy code
$ go run main.go
This command will download any necessary dependencies and start the server.

Access the forum in your web browser:
Copy code
http://localhost:8000
You should now see the For'Rhum Arrangé homepage.

Functionality
The For'Rhum Arrangé provides the following features:

Categories
To view the different categories, click on the forum logo.
The categories will be displayed, allowing you to explore and navigate to specific topics of interest.
Creating a Post
To create a new post, navigate to the forum homepage.
Fill in the required details, including the post title, content, and optional images.
Click "Submit" to publish your post to the forum.
User Account
To create a user account, click on the logo profile button on the forum homepage.
Fill in the necessary details, such as username, email, and password.
Click "Submit" to create your account.
To log in to your account, provide your credentials.
Database
The forum application is connected to a database to store and retrieve data relating to user accounts. The necessary database setup and configuration have been handled within the application.