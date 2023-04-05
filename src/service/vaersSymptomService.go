package service

type VaersSymptomService interface {
	GetVaersSymptomListByVaersId(id int64) (list []*VaersSymptom, err error)
	GetVaersSymptomTermList(keyword string, page, pageSize int) (list []*VaersSymptomTerm, count int64, err error)
	GetVaersSymptomTerm(id int64) (term *VaersSymptomTerm, err error)
	CountSymptomByVaersId(id int64) (count int64, err error)
	CountSymptom() (count int64, err error)
	GetVaersIdListBySymptomId(id int64) (list []int64, err error)
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
