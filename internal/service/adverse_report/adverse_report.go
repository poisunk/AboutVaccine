package adverse_report

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/vaccine"
	"sync"
)

type AdverseEventRepo interface {
	Create(event *entity.AdverseEvent) error
	Get(id int64) (*entity.AdverseEvent, bool, error)
	GetList(page, pageSize int) ([]*entity.AdverseEvent, int64, error)
	GetListByUid(uid int64, page, pageSize int) ([]*entity.AdverseEvent, int64, error)
	Delete(id int64) error
}

type AdverseSymptomRepo interface {
	CreateList(symptoms []*entity.AdverseSymptom) error
	GetListByEventId(eventId int64) ([]*entity.AdverseSymptom, error)
}

type AdverseVaccineRepo interface {
	CreateList(vaccines []*entity.AdverseVaccine) error
	GetListByEventId(eventId int64) ([]*entity.AdverseVaccine, error)
}

type AdverseReportCommon struct {
	adverseEventRepo   AdverseEventRepo
	adverseSymptomRepo AdverseSymptomRepo
	adverseVaccineRepo AdverseVaccineRepo
	vaccineCommon      *vaccine.VaccineCommon
}

func NewAdverseEventCommon(
	adverseEventRepo AdverseEventRepo,
	adverseSymptomRepo AdverseSymptomRepo,
	adverseVaccineRepo AdverseVaccineRepo,
	vaccineCommon *vaccine.VaccineCommon,
) *AdverseReportCommon {
	return &AdverseReportCommon{
		adverseEventRepo:   adverseEventRepo,
		adverseSymptomRepo: adverseSymptomRepo,
		adverseVaccineRepo: adverseVaccineRepo,
		vaccineCommon:      vaccineCommon,
	}
}

func (a *AdverseReportCommon) Get(id int64) (*schema.AdverseEventInfo, bool, error) {
	event := &entity.AdverseEvent{}
	event, has, err := a.adverseEventRepo.Get(id)
	if err != nil {
		return nil, has, err
	}
	vaccineList, err := a.GetVaccineListByEventId(id)
	if err != nil {
		return nil, has, err
	}

	symptomList, err := a.GetSymptomListByEventId(id)
	if err != nil {
		return nil, has, err
	}
	eventInfo := a.FormatEventInfo(event)
	eventInfo.VaccineList = vaccineList
	eventInfo.SymptomList = symptomList
	return eventInfo, has, nil
}

func (a *AdverseReportCommon) GetList(page, pageSize int) ([]*schema.AdverseEventInfo, int64, error) {
	events, total, err := a.adverseEventRepo.GetList(page, pageSize)
	if err != nil {
		return nil, total, err
	}
	eventInfos := make([]*schema.AdverseEventInfo, 0)
	var wg sync.WaitGroup
	for _, event := range events {
		eventInfo := a.FormatEventInfo(event)
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.LoadVaccineAndSymptomList(eventInfo)
		}()
		eventInfos = append(eventInfos, eventInfo)
	}
	wg.Wait()
	return eventInfos, total, nil
}

func (a *AdverseReportCommon) Create(schema *schema.AdverseEventAdd, uid *int64) error {
	event, vaccines, symptoms := a.FormatEntity(schema, uid)
	err := a.adverseEventRepo.Create(event)
	if err != nil {
		return err
	}
	for _, s := range symptoms {
		s.EventId = event.Id
	}
	err = a.adverseSymptomRepo.CreateList(symptoms)
	if err != nil {
		return err
	}
	for _, v := range vaccines {
		v.EventId = event.Id
	}
	err = a.adverseVaccineRepo.CreateList(vaccines)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdverseReportCommon) GetVaccineListByEventId(id int64) ([]*schema.AdverseVaccineInfo, error) {
	vaccineEntitys, err := a.adverseVaccineRepo.GetListByEventId(id)
	if err != nil {
		return nil, err
	}
	vaccineList := make([]*schema.AdverseVaccineInfo, 0)
	for _, v := range vaccineEntitys {
		vaccineInfo := a.FormatVaccineInfo(v, func(info *schema.AdverseVaccineInfo) {
			example, _, err := a.vaccineCommon.Get(v.VaccineId)
			if err != nil {
				return
			}
			info.Type = example.Type
			info.Name = example.ProductName
			info.Manufacturer = example.ProductionCompany
		})
		vaccineList = append(vaccineList, vaccineInfo)
	}
	return vaccineList, nil
}

func (a *AdverseReportCommon) GetSymptomListByEventId(id int64) ([]*schema.AdverseSymptomInfo, error) {
	symptomEntitys, err := a.adverseSymptomRepo.GetListByEventId(id)
	if err != nil {
		return nil, err
	}
	symptomList := make([]*schema.AdverseSymptomInfo, 0)
	for _, s := range symptomEntitys {
		symptom := a.FormatSymptomInfo(s, nil)
		symptomList = append(symptomList, symptom)
	}
	return symptomList, nil
}

func (a *AdverseReportCommon) GetListByUid(uid int64, page, pageSize int) ([]*schema.AdverseEventInfo, int64, error) {
	events, total, err := a.adverseEventRepo.GetListByUid(uid, page, pageSize)
	if err != nil {
		return nil, total, err
	}
	eventInfos := make([]*schema.AdverseEventInfo, 0)
	var wg sync.WaitGroup
	for _, event := range events {
		eventInfo := a.FormatEventInfo(event)
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.LoadVaccineAndSymptomList(eventInfo)
		}()
		eventInfos = append(eventInfos, eventInfo)
	}
	wg.Wait()
	return eventInfos, total, nil
}

func (a *AdverseReportCommon) LoadVaccineAndSymptomList(event *schema.AdverseEventInfo) error {
	vaccineList, err := a.GetVaccineListByEventId(event.Id)
	if err != nil {
		return err
	}
	symptomList, err := a.GetSymptomListByEventId(event.Id)
	if err != nil {
		return err
	}
	event.VaccineList = vaccineList
	event.SymptomList = symptomList
	return nil
}

func (a *AdverseReportCommon) Delete(id int64) error {
	return a.adverseEventRepo.Delete(id)
}
