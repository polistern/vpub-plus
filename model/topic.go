package model

import "time"

type Topic struct {
	Id          int64
	BoardId     int64
	Subject     string
	Author      string
	FirstPostId int64
	Replies     int64
	UpdatedAt   time.Time
}
