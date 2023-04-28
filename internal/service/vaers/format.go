package vaers

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

func (vc *VaersResultCommon) FormatResultInfo(v *entity.VaersResult) *schema.VaersResultInfo {
	return &schema.VaersResultInfo{
		VaccineId: v.VaccineId,
		SymptomId: v.SymptomId,
		Vaccine:   v.Name,
		Symptom:   v.Symptom,
		Total:     v.Total,
	}
}

func (vx *VaersVaxCommon) FormatVaxTerm(v *entity.VaersVaxTerm) *schema.VaersVaxTerm {
	return &schema.VaersVaxTerm{
		Id:           v.Id,
		Type:         v.Type,
		Manufacturer: v.Manufacturer,
		Name:         v.Name,
	}
}

func (vs *VaersSymptomCommon) FormatSymptomTerm(v *entity.VaersSymptomTerm) *schema.VaersSymptomTerm {
	return &schema.VaersSymptomTerm{
		Id:      v.Id,
		Symptom: v.Symptom,
	}
}
