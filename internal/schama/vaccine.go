package schama

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/utile"
)

type Vaccine struct {
	Id                int64  `json:"id"`
	Type              string `json:"type"`
	RegisterNumber    string `json:"registerNumber"`
	ProductName       string `json:"productName"`
	EnglishName       string `json:"englishName"`
	TradeName         string `json:"tradeName"`
	Dosage            string `json:"dosage"`
	Specification     string `json:"specification"`
	Owner             string `json:"owner"`
	OwnerAddress      string `json:"ownerAddress"`
	ProductionCompany string `json:"productionCompany"`
	ApprovalDate      string `json:"approvalDate"`
	ProductionAddress string `json:"productionAddress"`
	ProductionClass   string `json:"productionClass"`
	OriginalNumber    string `json:"originalNumber"`
	DrugCode          string `json:"drugCode"`
	DrugCodeNote      string `json:"drugCodeNote"`
}

type VaccineType struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func (v *Vaccine) GetFormEntity(e *entity.Vaccine) {
	_ = utile.StructConv(e, v)
}

func (v *VaccineType) GetFormEntity(e *entity.VaccineType) {
	_ = utile.StructConv(e, v)
}
