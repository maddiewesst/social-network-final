# Database Schema

This document outlines the database schema for the Social Network Project. The database is implemented using SQLite and adheres to an entity-relationship model to efficiently manage data.

## Tables

### Users
Stores user information and profile details.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the user.     |
| email          | TEXT       | UNIQUE NOT NULL       | User's email address.               |
| password       | TEXT       | NOT NULL              | Hashed password for authentication. |
| first_name     | TEXT       | NOT NULL              | User's first name.                  |
| last_name      | TEXT       | NOT NULL              | User's last name.                   |
| date_of_birth  | DATE       | NOT NULL              | User's date of birth.               |
| avatar         | TEXT       | NULL                  | Path to user's avatar image.        |
| nickname       | TEXT       | NULL                  | Optional nickname.                  |
| about_me       | TEXT       | NULL                  | Optional profile description.       |
| is_private     | BOOLEAN    | DEFAULT FALSE         | Indicates if the profile is private.|

---

### Posts
Stores user posts with privacy settings.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the post.     |
| user_id        | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the post creator.            |
| content        | TEXT       | NOT NULL              | Text content of the post.           |
| image          | TEXT       | NULL                  | Path to attached image/GIF.         |
| privacy        | TEXT       | CHECK(privacy IN ('public', 'private', 'restricted')) | Privacy level of the post. |
| created_at     | TIMESTAMP  | DEFAULT CURRENT_TIMESTAMP | Timestamp of post creation.         |

---

### Followers
Manages follower relationships between users.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the relationship. |
| follower_id    | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | User following another user.        |
| followed_id    | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | User being followed.                |
| status         | TEXT       | CHECK(status IN ('pending', 'accepted')) | Status of the follow request.       |

---

### Groups
Stores group details.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the group.    |
| creator_id     | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the group creator.            |
| title          | TEXT       | NOT NULL              | Title of the group.                 |
| description    | TEXT       | NULL                  | Group description.                  |

---

### Group Members
Manages group membership.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the membership.|
| group_id       | INTEGER    | FOREIGN KEY REFERENCES Groups(id) ON DELETE CASCADE | ID of the group.                    |
| user_id        | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the user.                     |
| role           | TEXT       | CHECK(role IN ('member', 'admin')) | Role of the user in the group.      |

---

### Notifications
Stores notifications for users.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the notification.|
| user_id        | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the recipient user.           |
| content        | TEXT       | NOT NULL              | Notification message.               |
| is_read        | BOOLEAN    | DEFAULT FALSE         | Indicates if the notification is read.|
| created_at     | TIMESTAMP  | DEFAULT CURRENT_TIMESTAMP | Timestamp of notification creation. |

---

### Chats
Manages private messages between users.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the message.  |
| sender_id      | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the message sender.           |
| recipient_id   | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the message recipient.        |
| content        | TEXT       | NOT NULL              | Text content of the message.        |
| created_at     | TIMESTAMP  | DEFAULT CURRENT_TIMESTAMP | Timestamp of message creation.      |

---

### Group Chats
Manages messages in group chat rooms.

| Column         | Type       | Constraints           | Description                          |
|----------------|------------|-----------------------|--------------------------------------|
| id             | INTEGER    | PRIMARY KEY AUTOINCREMENT | Unique identifier for the message.  |
| group_id       | INTEGER    | FOREIGN KEY REFERENCES Groups(id) ON DELETE CASCADE | ID of the group.                    |
| sender_id      | INTEGER    | FOREIGN KEY REFERENCES Users(id) ON DELETE CASCADE | ID of the message sender.           |
| content        | TEXT       | NOT NULL              | Text content of the message.        |
| created_at     | TIMESTAMP  | DEFAULT CURRENT_TIMESTAMP | Timestamp of message creation.      |

---

## Entity-Relationship Diagram (ERD)
For a visual representation of the database schema, refer to the `erd.png` file in the `docs/` folder.


