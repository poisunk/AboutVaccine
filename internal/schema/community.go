package schema

import (
	"time"
	"vax/internal/entity"
)

type Issue struct {
	Id          int64     `json:"id"`
	OwnerId     int64     `json:"ownerId"`
	Owner       string    `json:"owner"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateOn    time.Time `json:"createOn"`
	UpdateOn    time.Time `json:"updateOn"`
}

type Answer struct {
	Id       int64     `json:"id"`
	OwnerId  int64     `json:"ownerId"`
	Owner    string    `json:"owner"`
	IssueId  int64     `json:"issueId"`
	Content  string    `json:"content"`
	CreateOn time.Time `json:"createOn"`
	UpdateOn time.Time `json:"updateOn"`
}

func (i *Issue) ToEntity() *entity.Issue {
	return &entity.Issue{
		Uid:         i.OwnerId,
		Title:       i.Title,
		Description: i.Description,
	}
}

func (i *Issue) GetFormEntity(e *entity.Issue) {
	i.Id = e.Id
	i.OwnerId = e.Uid
	i.Title = e.Title
	i.Description = e.Description
	i.CreateOn = e.CreateOn
	i.UpdateOn = e.UpdateOn
}
