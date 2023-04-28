package schema

type VaersVaxInfo struct {
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
