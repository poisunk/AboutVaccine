package adverse_report

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/user"
	"about-vaccine/internal/service/vaccine"
	"sync"
)

type AdverseEventRepo interface {
	Create(event *entity.AdverseEvent) error
	Get(id int64) (*entity.AdverseEvent, bool, error)
	GetBriefList(page, pageSize int) ([]*entity.AdverseEvent, int64, error)
	GetBriefListByUid(uid int64, page, pageSize int) ([]*entity.AdverseEvent, int64, error)
	GetBriefListByKeyword(keyword string, page, pageSize int) ([]*entity.AdverseEvent, int64, error)
	GetBriefListByVaccineId(vid int64, page, pageSize int) ([]*entity.AdverseEvent, error)
	GetBriefListByOAEId(oid int64, page, pageSize int) ([]*entity.AdverseEvent, error)
	GetUid(id int64) (int64, bool, error)
	Delete(id int64) error
	Count() (int64, error)
	CountByVaccineId(vid int64) (int64, error)
	CountByOAEId(oid int64) (int64, error)
}

type AdverseSymptomRepo interface {
	CreateList(symptoms []*entity.AdverseSymptom) error
	GetListByEventId(eventId int64) ([]*entity.AdverseSymptom, error)
	GetListByVaccineId(vid int64, page, pageSize int) ([]*entity.AdverseSymptom, error)
}

type AdverseVaccineRepo interface {
	CreateList(vaccines []*entity.AdverseVaccine) error
	GetListByEventId(eventId int64) ([]*entity.AdverseVaccine, error)
	GetListByOAEId(oid int64, page, pageSize int) ([]*entity.AdverseVaccine, error)
}

type AdverseReportCommon struct {
	adverseEventRepo   AdverseEventRepo
	adverseSymptomRepo AdverseSymptomRepo
	adverseVaccineRepo AdverseVaccineRepo
	vaccineCommon      *vaccine.VaccineCommon
	userCommon         *user.UserCommon
}

func NewAdverseEventCommon(
	adverseEventRepo AdverseEventRepo,
	adverseSymptomRepo AdverseSymptomRepo,
	adverseVaccineRepo AdverseVaccineRepo,
	vaccineCommon *vaccine.VaccineCommon,
	userCommon *user.UserCommon,
) *AdverseReportCommon {
	return &AdverseReportCommon{
		adverseEventRepo:   adverseEventRepo,
		adverseSymptomRepo: adverseSymptomRepo,
		adverseVaccineRepo: adverseVaccineRepo,
		vaccineCommon:      vaccineCommon,
		userCommon:         userCommon,
	}
}

func (a *AdverseReportCommon) Get(id int64) (*schema.AdverseEventInfo, bool, error) {
	event := &entity.AdverseEvent{}
	event, has, err := a.adverseEventRepo.Get(id)
	if err != nil {
		return nil, has, err
	}
	eventInfo := a.FormatEventInfo(event)
	a.LoadVaccineAndSymptomList(eventInfo)
	eventInfo.UserName = a.LoadUserName(event.Uid)
	return eventInfo, has, nil
}

func (a *AdverseReportCommon) GetList(page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	events, total, err := a.adverseEventRepo.GetBriefList(page, pageSize)
	if err != nil {
		return nil, total, err
	}
	eventInfos := make([]*schema.AdverseEventBriefInfo, 0)
	var wg sync.WaitGroup
	for _, event := range events {
		eventInfo := a.FormatEventBriefInfo(event)
		wg.Add(1)
		go func() {
			defer wg.Done()
			symptomList, _ := a.GetSymptomListByEventId(event.Id)
			eventInfo.SymptomList = symptomList
			eventInfo.UserName = a.LoadUserName(event.Uid)
		}()
		eventInfos = append(eventInfos, eventInfo)
	}
	wg.Wait()
	return eventInfos, total, nil
}

func (a *AdverseReportCommon) GetListByKeyword(keyword string, page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	events, total, err := a.adverseEventRepo.GetBriefListByKeyword(keyword, page, pageSize)
	if err != nil {
		return nil, total, err
	}
	eventInfos := make([]*schema.AdverseEventBriefInfo, 0)
	var wg sync.WaitGroup
	for _, event := range events {
		eventInfo := a.FormatEventBriefInfo(event)
		wg.Add(1)
		go func() {
			defer wg.Done()
			symptomList, _ := a.GetSymptomListByEventId(event.Id)
			eventInfo.SymptomList = symptomList
			eventInfo.UserName = a.LoadUserName(event.Uid)
		}()
		eventInfos = append(eventInfos, eventInfo)
	}
	wg.Wait()
	return eventInfos, total, nil
}

func (a *AdverseReportCommon) GetListByUid(uid int64, page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	events, total, err := a.adverseEventRepo.GetBriefListByUid(uid, page, pageSize)
	if err != nil {
		return nil, total, err
	}
	eventInfos := make([]*schema.AdverseEventBriefInfo, 0)
	var wg sync.WaitGroup
	for _, event := range events {
		eventInfo := a.FormatEventBriefInfo(event)
		wg.Add(1)
		go func(e *entity.AdverseEvent) {
			defer wg.Done()
			symptomList, _ := a.GetSymptomListByEventId(e.Id)
			eventInfo.SymptomList = symptomList
			eventInfo.UserName = a.LoadUserName(e.Uid)
		}(event)
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
	var wg sync.WaitGroup
	for _, v := range vaccineEntitys {
		wg.Add(1)
		go func() {
			defer wg.Done()
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
		}()
	}
	wg.Wait()
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

func (a *AdverseReportCommon) GetListByVaccineId(id int64, page, pageSize int) ([]*schema.AdverseEventBriefInfo, error) {
	entitys, err := a.adverseEventRepo.GetBriefListByVaccineId(id, page, pageSize)
	if err != nil {
		return nil, err
	}
	eventInfos := make([]*schema.AdverseEventBriefInfo, 0, len(entitys))
	var wg sync.WaitGroup
	for _, v := range entitys {
		wg.Add(1)
		go func(e *entity.AdverseEvent) {
			defer wg.Done()
			vaccineInfo := a.FormatEventBriefInfo(e)
			eventInfos = append(eventInfos, vaccineInfo)
		}(v)
	}
	wg.Wait()
	return eventInfos, nil
}

func (a *AdverseReportCommon) GetListByOAEId(oid int64, page, pageSize int) ([]*schema.AdverseEventBriefInfo, error) {
	entitys, err := a.adverseEventRepo.GetBriefListByOAEId(oid, page, pageSize)
	if err != nil {
		return nil, err
	}
	eventInfos := make([]*schema.AdverseEventBriefInfo, 0, len(entitys))
	var wg sync.WaitGroup
	for _, v := range entitys {
		wg.Add(1)
		go func(e *entity.AdverseEvent) {
			defer wg.Done()
			vaccineInfo := a.FormatEventBriefInfo(e)
			eventInfos = append(eventInfos, vaccineInfo)
		}(v)
	}
	wg.Wait()
	return eventInfos, nil
}

func (a *AdverseReportCommon) GetUid(id int64) (int64, bool, error) {
	return a.adverseEventRepo.GetUid(id)
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

func (a *AdverseReportCommon) LoadUserName(uid *int64) string {
	if uid == nil {
		return ""
	}
	name, err := a.userCommon.GetUserNameByUid(*uid)
	if err != nil {
		return ""
	}
	return name
}

func (a *AdverseReportCommon) Delete(id int64) error {
	return a.adverseEventRepo.Delete(id)
}
