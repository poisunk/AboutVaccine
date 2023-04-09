package service

type VaccineService interface {
	CreateVaccine(Vaccine) error
	GetVaccine(int64) (*Vaccine, error)
	GetVaccineList(int, int, string) ([]*Vaccine, int, error)
	DeleteVaccine(int64) error
	GetVaccineExampleList(int, int, int) ([]*VaccineExample, int, error)
	GetVaccineExampleByTid(int64, int) (*VaccineExample, error)
}

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

type VaccineExample struct {
	Type        string     `json:"type"`
	VaccineList []*Vaccine `json:"vaccineList"`
	More        bool       `json:"more"`
}
