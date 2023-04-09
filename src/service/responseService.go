package service

import "time"

type ResponseService interface {
	GetResponsesByQid(qid int64, page int, pageSize int) (r []*Response, total int, err error)
	GetResponseByUid(uid int64, page int, pageSize int) ([]*Response, int, error)
	DeleteResponse(id int64) (err error)
	CreateResponse(r *Response) (err error)
	GetResponseTermsByRid(rid int64) (r []*ResponseTerm, err error)
	GetResponseOwnerId(rid int64) (uid int64, err error)
}

type Response struct {
	Id               int64           `json:"id"`
	QuestionnaireId  int64           `json:"questionnaireId"`
	UserId           int64           `json:"userId"`
	UserName         string          `json:"userName"`
	ResponseTime     time.Time       `gorm:"autoCreateTime;datetime;column:response_time" json:"responseTime"`
	ResponseTermList []*ResponseTerm `json:"responseTermList"`
}

type ResponseTerm struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"questionId"`
	Answer     string `json:"answer"`
}
