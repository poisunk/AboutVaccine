package entity

type OAETerm struct {
	Id              int    `json:"id"`
	TermIRI         string `json:"termIRI"`
	TermLabel       string `json:"termLabel"`
	ParentTermIRI   string `json:"parentTermIRI"`
	ParentTermLabel string `json:"parentTermLabel"`
	AlternativeTerm string `json:"alternativeTerm"`
	Definition      string `json:"definition"`
}

func (o *OAETerm) TableName() string {
	return "oae_term"
}
