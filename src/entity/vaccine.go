package entity

type Vaccine struct {
	Id                int64  `json:"id"`
	Tid               int64  `json:"tid"`
	RegisterNumber    string `gorm:"varchar(255);column:批准文号" json:"registerNumber"`
	ProductName       string `gorm:"varchar(255);column:产品名称" json:"productName"`
	EnglishName       string `gorm:"varchar(255);column:英文名称" json:"englishName"`
	TradeName         string `gorm:"varchar(255);column:商品名" json:"tradeName"`
	Dosage            string `gorm:"varchar(255);column:剂型" json:"dosage"`
	Specification     string `gorm:"varchar(255);column:规格" json:"specification"`
	Owner             string `gorm:"varchar(255);column:上市许可持有人" json:"owner"`
	OwnerAddress      string `gorm:"varchar(255);column:上市许可持有人地址" json:"ownerAddress"`
	ProductionCompany string `gorm:"varchar(255);column:生产单位" json:"productionCompany"`
	ApprovalDate      string `gorm:"varchar(255);column:批准日期" json:"approvalDate"`
	ProductionAddress string `gorm:"varchar(255);column:生产地址" json:"productionAddress"`
	ProductionClass   string `gorm:"varchar(255);column:产品类别" json:"productionClass"`
	OriginalNumber    string `gorm:"varchar(255);column:原批准文号" json:"originalNumber"`
	DrugCode          string `gorm:"varchar(255);column:药品本位码" json:"drugCode"`
	DrugCodeNote      string `gorm:"varchar(255);column:药品本位码备注" json:"drugCodeNote"`
}

func (v *Vaccine) TableName() string {
	return "vaccine_cfda"
}
