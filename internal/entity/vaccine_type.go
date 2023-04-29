package entity

type VaccineType struct {
	Id                  int64  `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	Type                string `xorm:"notnull VARCHAR(255) type" json:"type"`
	DiseaseIntroduction string `xorm:"notnull TEXT disease_introduction" json:"disease_introduction"`
	PreventiveMeasures  string `xorm:"notnull TEXT preventive_measures" json:"preventive_measures"`
	Target              string `xorm:"notnull TEXT target" json:"target"`
	VaccinationBan      string `xorm:"notnull TEXT vaccination_ban" json:"vaccination_ban"`
	AdverseEvent        string `xorm:"notnull TEXT adverse_event" json:"adverse_event"`
}

func (t *VaccineType) TableName() string {
	return "vaccine_type"
}
