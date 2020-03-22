package pkg

import "time"

type PageInfo struct {
	Page       string     `json:"page"`
	PerPage    int32      `json:"per_page"`
	Total      int32      `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []UserInfo `json:"data"`
}

type UserInfo struct {
	ID              int       `json:"id"`
	UserName        string    `json:"username"`
	About           string    `json:"about"`
	Submitted       int32     `json:"submitted"`
	UpdatedAt       time.Time `json:"updated_at"`
	SubmissionCount int32     `json:"submission_count"`
	CommentCount    int32     `json:"comment_count"`
	CreatedAt       int32     `json:"created_at"`
}