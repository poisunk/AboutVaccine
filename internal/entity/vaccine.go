package entity

type Vaccine struct {
	Id                int64  `json:"id"`
	Tid               int64  `json:"tid"`
	Type              string `xorm:"varchar(255) type" json:"type"`
	RegisterNumber    string `xorm:"varchar(255) 批准文号" json:"registerNumber"`
	ProductName       string `xorm:"varchar(255) 产品名称" json:"productName"`
	EnglishName       string `xorm:"varchar(255) 英文名称" json:"englishName"`
	TradeName         string `xorm:"varchar(255) 商品名" json:"tradeName"`
	Dosage            string `xorm:"varchar(255) 剂型" json:"dosage"`
	Specification     string `xorm:"varchar(255) 规格" json:"specification"`
	Owner             string `xorm:"varchar(255) 上市许可持有人" json:"owner"`
	OwnerAddress      string `xorm:"varchar(255) 上市许可持有人地址" json:"ownerAddress"`
	ProductionCompany string `xorm:"varchar(255) 生产单位" json:"productionCompany"`
	ApprovalDate      string `xorm:"varchar(255) 批准日期" json:"approvalDate"`
	ProductionAddress string `xorm:"varchar(255) 生产地址" json:"productionAddress"`
	ProductionClass   string `xorm:"varchar(255) 产品类别" json:"productionClass"`
	OriginalNumber    string `xorm:"varchar(255) 原批准文号" json:"originalNumber"`
	DrugCode          string `xorm:"varchar(255) 药品本位码" json:"drugCode"`
	DrugCodeNote      string `xorm:"varchar(255) 药品本位码备注" json:"drugCodeNote"`
}

func (v *Vaccine) TableName() string {
	return "vaccine_cfda"
}
