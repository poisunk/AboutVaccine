package entity

type OAETerm struct {
	Id              int    `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	TermIRI         string `xorm:"TermIRI" json:"termIRI"`
	TermLabel       string `xorm:"TermLabel" json:"termLabel"`
	ParentTermIRI   string `xorm:"ParentTermIRI" json:"parentTermIRI"`
	ParentTermLabel string `xorm:"ParentTermLabel" json:"parentTermLabel"`
	AlternativeTerm string `xorm:"AlternativeTerm" json:"alternativeTerm"`
	Definition      string `xorm:"Definition" json:"definition"`
}

func (o *OAETerm) TableName() string {
	return "oae_term"
}
