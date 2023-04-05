package service

type VaersSymptomService interface {
	GetVaersSymptomListByVaersId(id int64) (list []*VaersSymptom, err error)
	GetVaersSymptomTermList(keyword string, page, pageSize int) (list []*VaersSymptomTerm, err error)
}

type VaersSymptom struct {
	ID        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

type VaersSymptomTerm struct {
	ID      int64  `json:"id"`
	Symptom string `json:"symptom"`
}
