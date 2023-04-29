package schema

type VaccineTypeInfo struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

type VaccineTypeDetailInfo struct {
	Id                  int64  `json:"id"`
	Type                string `json:"type"`
	DiseaseIntroduction string `json:"disease_introduction"`
	PreventiveMeasures  string `json:"preventive_measures"`
	Target              string `json:"target"`
	VaccinationBan      string `json:"vaccination_ban"`
	AdverseEvent        string `json:"adverse_event"`
}
