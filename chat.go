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
	Time  time.Time  `json:"time"`
}

var ChatHistory = make([]*Message, 0)