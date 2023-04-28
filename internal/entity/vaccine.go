package entity

type Vaccine struct {
	Id                int64  `json:"id"`
	Tid               int64  `json:"tid"`
	Type              string `xorm:"VARCHAR(255) type" json:"type"`
	RegisterNumber    string `xorm:"VARCHAR(255) register_number" json:"registerNumber"`
	ProductName       string `xorm:"VARCHAR(255) product_name" json:"productName"`
	EnglishName       string `xorm:"VARCHAR(255) english_name" json:"englishName"`
	TradeName         string `xorm:"VARCHAR(255) trade_name" json:"tradeName"`
	Dosage            string `xorm:"VARCHAR(255) dosage" json:"dosage"`
	Specification     string `xorm:"VARCHAR(255) specification" json:"specification"`
	Owner             string `xorm:"VARCHAR(255) owner" json:"owner"`
	OwnerAddress      string `xorm:"VARCHAR(255) owner_address" json:"ownerAddress"`
	ProductionCompany string `xorm:"VARCHAR(255) production_company" json:"productionCompany"`
	ApprovalDate      string `xorm:"VARCHAR(255) approval_date" json:"approvalDate"`
	ProductionAddress string `xorm:"VARCHAR(255) production_address" json:"productionAddress"`
	ProductionClass   string `xorm:"VARCHAR(255) production_class" json:"productionClass"`
	OriginalNumber    string `xorm:"VARCHAR(255) original_number" json:"originalNumber"`
	DrugCode          string `xorm:"VARCHAR(255) drug_code" json:"drugCode"`
	DrugCodeNote      string `xorm:"VARCHAR(255) drug_code_note" json:"drugCodeNote"`
}

func (v *Vaccine) TableName() string {
	return "vaccine_cfda"
}
