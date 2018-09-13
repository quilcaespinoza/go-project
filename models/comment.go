package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID   uint      `json:"userId"`
	ParentID uint      `json:"parentId"`
	Votes    uint32    `json:"votes"`
	Content  string    `json:"content"`
	HasVote  uint8     `json:"hasVote" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
