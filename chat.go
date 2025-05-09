package main

import (
	"time"
)

type Role string

const (
	// RoleUser is the role for a regular user
	User Role = "user"
	// Agent is the role for the AI
	Agent Role = "agent"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"message"`
	Time    string `json:"time"`
}

func NewMessage(role Role, content string) *Message {
	return &Message{
		Role:    role,
		Content: content,
		Time:    formatRelativeTime(time.Now()),
	}
}

var ChatHistory = make([]*Message, 0)

func formatRelativeTime(t time.Time) string {
	now := time.Now()

	timeStr := t.Format("3:04 PM")

	if now.YearDay() == t.YearDay() && now.Year() == t.Year() {
		return timeStr
	} else if now.YearDay() == t.YearDay()+1 && now.Year() == t.Year() {
		return "Yesterday " + timeStr
	} else {
		// For other days, show the date and time
		return t.Format("Jan 2") + " " + timeStr
	}
}