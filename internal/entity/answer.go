package entity

import "time"

type Answer struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`
	IssueId   int64     `json:"issueId"`
	Content   string    `json:"content"`
	CreateOn  time.Time `xorm:"created" json:"createOn"`
	UpdateOn  time.Time `xorm:"updated" json:"updateOn"`
	IsDeleted bool      `xorm:"deleted" json:"isDelete"`
}

func (a *Answer) TableName() string {
	return "answer"
}
