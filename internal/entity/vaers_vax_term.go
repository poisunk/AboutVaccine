package entity

type VaersVaxTerm struct {
	Id           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

func (v *VaersVaxTerm) TableName() string {
	return "vaers_vax_term"
}
