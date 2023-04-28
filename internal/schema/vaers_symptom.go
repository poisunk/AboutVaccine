package schema

type VaersSymptomInfo struct {
	Id        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

type VaersSymptomTerm struct {
	Id      int64  `json:"id"`
	Symptom string `json:"symptom"`
}
