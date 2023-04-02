package service

type OAEService interface {
	GetOAETermByLabel(label string, page int, pageSize int) (oaeList []*OAETerm, total int, err error)
	GetOAETermByIRI(IRI string) (oaeTerm *OAETerm, err error)
	GetOAETermParentList(IRI string) (oaeList []*OAETerm, err error)
}

type OAETerm struct {
	TermIRI         string `json:"termIRI"`
	TermLabel       string `json:"termLabel"`
	ParentTermIRI   string `json:"parentTermIRI"`
	ParentTermLabel string `json:"parentTermLabel"`
	AlternativeTerm string `json:"alternativeTerm"`
	Definition      string `json:"definition"`
}
