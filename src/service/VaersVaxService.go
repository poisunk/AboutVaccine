package service

type VaersVaxService interface {
	GetVaersVaxListByVaersId(vaersId int64) (list []*VaersVax, err error)
	GetVaersVaxTermList(keyword string, page, pageSize int) (list []*VaersVaxTerm, err error)
}

type VaersVax struct {
	ID           int64  `json:"id"`
	VaersId      int64  `json:"vaersId"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Dose         string `json:"dose"`
	Route        string `json:"route"`
	Site         string `json:"site"`
	VaxId        int64  `json:"vaxId"`
}

type VaersVaxTerm struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}
