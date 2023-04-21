package entity

type VaccineType struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func (t *VaccineType) TableName() string {
	return "vaccine_type"
}
