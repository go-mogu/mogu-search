package model

import "time"

type BlogSort struct {
	Uid        string    `gorm:"primaryKey" json:"uid"`
	SortName   string    `json:"sortName"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt  time.Time `gorm:"column:update_time" json:"updateTime"`
	Status     int8      `gorm:"default:1" json:"status"`
	Sort       int       `json:"sort"`
	ClickCount int       `json:"clickCount"`
}
