// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package crud

import (
	"time"
)

type Group struct {
	ID          int64
	Title       string
	Creator     int64
	Description string
	CreatedAt   time.Time
}

type GroupChatItem struct {
	ID        int64
	GroupID   int64
	UserID    int64
	ChatNoti  int64
	LastMsgAt time.Time
}

type GroupEvent struct {
	ID          int64
	Author      int64
	GroupID     int64
	Title       string
	Description string
	CreatedAt   time.Time
	Date        time.Time
}

type GroupEventMember struct {
	ID      int64
	UserID  int64
	EventID int64
	Status  int64
}

type GroupMember struct {
	ID       int64
	UserID   int64
	GroupID  int64
	Status   int64
	ChatNoti int64
}

type GroupMessage struct {
	ID        int64
	SourceID  int64
	GroupID   int64
	Message   string
	CreatedAt time.Time
}

type GroupPost struct {
	ID        int64
	Author    int64
	GroupID   int64
	Message   string
	Image     string
	CreatedAt time.Time
}

type GroupPostComment struct {
	ID          int64
	Author      int64
	GroupPostID int64
	Message     string
	CreatedAt   time.Time
}

type GroupRequest struct {
	ID      int64
	UserID  int64
	GroupID int64
	Status  string
}

type Post struct {
	ID        int64
	Author    int64
	Message   string
	Image     string
	CreatedAt time.Time
	Privacy   int64
}

type PostComment struct {
	ID        int64
	UserID    int64
	PostID    int64
	CreatedAt time.Time
	Message   string
	Image     string
}

type PostMember struct {
	ID     int64
	UserID int64
	PostID int64
}

type PrivateChatItem struct {
	ID        int64
	SourceID  int64
	TargetID  int64
	ChatNoti  int64
	LastMsgAt time.Time
}

type SessionTable struct {
	SessionToken string
	UserID       int64
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	NickName  string
	Email     string
	Password  string
	Dob       time.Time
	Image     string
	About     string
	Public    int64
}

type UserFollower struct {
	ID       int64
	SourceID int64
	TargetID int64
	Status   int64
}

type UserMessage struct {
	ID        int64
	SourceID  int64
	TargetID  int64
	Message   string
	CreatedAt time.Time
}
