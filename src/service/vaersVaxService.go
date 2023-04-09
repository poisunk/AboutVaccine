package service

type VaersVaxService interface {
	GetVaersVaxListByVaersId(vaersId int64) (list []*VaersVax, err error)

	GetVaersVaxTerm(id int64) (term *VaersVaxTerm, err error)
	GetVaersVaxTermList(keyword string, page, pageSize int) (list []*VaersVaxTerm, total int64, err error)

	GetVaersIdListByVaxId(vaxId int64) (list []int64, err error)

	CountVaersVaxByVaersId(vaersId int64) (count int64, err error)
}

type VaersVax struct {
	Id           int64  `json:"id"`
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
	Id           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}
