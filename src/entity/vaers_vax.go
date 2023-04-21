package entity

type VaersVax struct {
	Id      int64  `json:"id"`
	VaersId int64  `json:"vaersId"`
	Dose    string `json:"dose"`
	Route   string `json:"route"`
	Site    string `json:"site"`
	VaxId   int64  `json:"vaxId"`
}

func (v *VaersVax) TableName() string {
	return "vaers_vax"
}
