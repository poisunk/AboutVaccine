package service

import (
	"about-vaccine/src/entity"
	"about-vaccine/src/repo"
	"about-vaccine/src/schama"
	"errors"
	"log"
)

type AdverseEventService struct {
	EventRepo      *repo.AdverseEventRepo
	SymptomRepo    *repo.AdverseSymptomRepo
	VaccineRepo    *repo.AdverseVaccineRepo
	VaccineService *VaccineService
}

func NewAdverseEventService(
	eventRepo *repo.AdverseEventRepo,
	symptomRepo *repo.AdverseSymptomRepo,
	vaccineRepo *repo.AdverseVaccineRepo,
) *AdverseEventService {
	return &AdverseEventService{
		EventRepo:   eventRepo,
		SymptomRepo: symptomRepo,
		VaccineRepo: vaccineRepo,
	}
}

func (a *AdverseEventService) Create(event *schama.AdverseEvent) error {
	// 1. 检查数据
	if len(event.Description) == 0 {
		return errors.New("描述不能为空")
	}
	// 2. 创建Event
	entityEvent := event.ToEntity()
	err := a.EventRepo.Create(entityEvent)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建Event失败")
	}
	// 3. 创建Vaccine
	var entityVaccineList []*entity.AdverseVaccine
	for _, vaccine := range event.VaccineList {
		entityVaccine := vaccine.ToEntity()
		entityVaccineList = append(entityVaccineList, entityVaccine)
	}
	err = a.VaccineRepo.Create(entityVaccineList...)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建Vaccine失败")
	}
	// 4. 创建Symptom
	var entitySymptomList []*entity.AdverseSymptom
	for _, symptom := range event.SymptomList {
		entitySymptom := symptom.ToEntity()
		entitySymptomList = append(entitySymptomList, entitySymptom)
	}
	err = a.SymptomRepo.Create(entitySymptomList...)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建Symptom失败")
	}
	return nil
}

func (a *AdverseEventService) Get(id int64) (*schama.AdverseEvent, error) {
	e, _, err := a.EventRepo.GetById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取Event失败")
	}
	// 组成Event
	event := &schama.AdverseEvent{}
	event.GetFromEntity(e)
	err = a.LoadRequisite(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (a *AdverseEventService) GetList(page, pageSize int) ([]*schama.AdverseEvent, int64, error) {
	el, total, err := a.EventRepo.GetList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	// 获取Vaccine，Symptom
	eventList := make([]*schama.AdverseEvent, len(el))
	for _, e := range el {
		event := &schama.AdverseEvent{}
		event.GetFromEntity(e)
		err := a.LoadRequisite(event)
		if err != nil {
			return nil, 0, err
		}
		eventList = append(eventList, event)
	}
	return eventList, total, nil
}

func (a *AdverseEventService) GetListByUid(uid int64, page, pageSize int) ([]*schama.AdverseEvent, int64, error) {
	el, total, err := a.EventRepo.GetListByUid(uid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	// 获取Vaccine，Symptom
	eventList := make([]*schama.AdverseEvent, len(el))
	for _, e := range el {
		event := &schama.AdverseEvent{}
		event.GetFromEntity(e)
		err := a.LoadRequisite(event)
		if err != nil {
			return nil, 0, err
		}
		eventList = append(eventList, event)
	}
	return eventList, total, nil
}

func (a *AdverseEventService) Delete(id int64) error {
	err := a.EventRepo.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除Event失败")
	}
	return nil
}

func (a *AdverseEventService) LoadRequisite(e *schama.AdverseEvent) error {
	// 获取Vaccine，组装VaccineList数据
	vl, err := a.VaccineRepo.GetByEventId(e.Id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("获取Vaccine列表失败")
	}
	vaccineList := make([]*schama.AdverseVaccine, len(vl))
	for _, v := range vl {
		vaccine := &schama.AdverseVaccine{}
		vax, err := a.VaccineService.Get(v.Id)
		if err != nil {
			log.Println(err.Error())
			return errors.New("获取Vaccine失败")
		}
		vaccine.GetFromVaccine(v, vax)
		vaccineList = append(vaccineList, vaccine)
	}
	// 获取Symptom，组装SymptomList数据
	sl, err := a.SymptomRepo.GetByEventId(e.Id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("获取Symptom列表失败")
	}
	symptomList := make([]*schama.AdverseSymptom, len(sl))
	for _, s := range sl {
		symptom := &schama.AdverseSymptom{}
		symptom.GetFromSymptom(s)
		symptomList = append(symptomList, symptom)
	}
	// 组成Event
	e.VaccineList = vaccineList
	e.SymptomList = symptomList
	return nil
}
