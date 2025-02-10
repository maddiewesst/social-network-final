# Social Network Project

## Overview
This project is a Facebook-like social network application that allows users to connect, share, and communicate. It includes features such as user authentication, profile management, post creation, groups, notifications, and real-time chats. 

## Features
The application includes the following features:

- User registration and login with sessions and cookies for authentication
- User profiles with public and private settings
- Ability to follow/unfollow other users
- Creation of posts and comments with privacy settings
- Creation of groups with invitations and requests
- Real-time private and group messaging using Websocket
- Notifications for following requests, group invitations and requests, and group events using Websocket
- Dockerized deployment with separate backend and frontend images

## Technologies Used

- Frontend: React 
- Backend: Golang 
- Database: SQLite
- WebSocket: Gorilla WebSocket package
- Migration: golang-migrate package
- Authentication: Sessions and cookies
- Containerization: Docker

## Setup Instructions

### Prerequisites
- Docker and Docker Compose installed.
- Go installed (for local development).
- Node.js installed (for frontend development).

### Installation
1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. **Setup the backend**:
   - Navigate to the backend folder:
     ```bash
     cd backend
     ```
   - Build the Docker image:
     ```bash
     docker build -t social-network-backend .
     ```

3. **Setup the frontend**:
   - Navigate to the frontend folder:
     ```bash
     cd frontend
     ```
   - Build the Docker image:
     ```bash
     docker build -t social-network-frontend .
     ```

4. **Run the application**:
   - Start the containers using Docker Compose:
     ```bash
     docker-compose up
     ```
   - Access the application at `http://localhost:3000`.

### Database Migrations
The project uses SQLC for database interactions and golang-migrate for migrations.
- Migration files are stored in `backend/pkg/db/migrations/sqlite`.
- To apply migrations manually:
  ```bash
  migrate -path backend/pkg/db/migrations/sqlite -database sqlite3://path/to/database up
  ```


![alt text](<Screenshot 2025-02-10 at 12.01.19.png>)