package models

import (
	"MyWeb/dao"
	"MyWeb/utile"
)

type Vaccine struct {
	ID                int64  `gorm:"int(10);column:id;primary_key" json:"id"`
	Tid               int64  `gson:"int(11);column:tid" json:"tid"`
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

func CreateVaccine(v *Vaccine) (err error) {
	err = dao.DB.Create(v).Error
	return
}

func GetVaccineList(page int, pageSize int, keyword string) (vList []*Vaccine, total int, err error) {
	db := dao.DB.Model(Vaccine{}).Where("`产品名称` LIKE ?", utile.HandleSearchWord(keyword)).Count(&total)
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&vList).Error
	return vList, total, err
}

func GetVaccineById(id int64) (v *Vaccine, err error) {
	v = &Vaccine{}
	if err = dao.DB.Where("id = ?", id).First(v).Error; err != nil {
		return nil, err
	}
	return
}

func GetVaccineListByTid(tid int64, limit int) (vList []*Vaccine, err error) {
	if err = dao.DB.Where("tid = ?", tid).Limit(limit).Find(&vList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteVaccine(id int64) (err error) {
	if err = dao.DB.Delete(Vaccine{}, "id = ?", id).Error; err != nil {
		return err
	}
	return
}
