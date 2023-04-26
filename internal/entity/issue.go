package entity

import "time"

type Issue struct {
	Id          int64     `json:"id"`
	Uid         int64     `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateOn    time.Time `xorm:"created" json:"createOn"`
	UpdateOn    time.Time `xorm:"updated" json:"updateOn"`
	IsDeleted   bool      `xorm:"deleted" json:"isDelete"`
}

func (i *Issue) TableName() string {
	return "issue"
}
