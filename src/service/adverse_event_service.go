package service

import (
	"about-vaccine/src/entity"
	"about-vaccine/src/repo"
	"about-vaccine/src/schama"
	"errors"
	"log"
	"sync"
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
	vaccineService *VaccineService,
) *AdverseEventService {
	return &AdverseEventService{
		EventRepo:      eventRepo,
		SymptomRepo:    symptomRepo,
		VaccineRepo:    vaccineRepo,
		VaccineService: vaccineService,
	}
}

func (a *AdverseEventService) Create(event *schama.AdverseEvent) error {
	// 1. 检查数据
	if len(event.Description) == 0 {
		return errors.New("描述不能为空")
	}
	// 2. 创建Event
	entityEvent := event.ToEntity()
	err := a.EventRepo.CreateOne(entityEvent)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建Event失败")
	}
	log.Println(entityEvent.Id)
	// 3. 创建Vaccine
	var entityVaccineList []*entity.AdverseVaccine
	for _, vaccine := range event.VaccineList {
		entityVaccine := vaccine.ToEntity()
		entityVaccine.AdverseEventId = entityEvent.Id
		entityVaccineList = append(entityVaccineList, entityVaccine)
	}
	err = a.VaccineRepo.Create(entityVaccineList...)
	if err != nil {
		log.Println(err.Error())
		go func() {
			_ = a.EventRepo.Delete(entityEvent.Id)
		}()
		return errors.New("创建Vaccine失败")
	}
	// 4. 创建Symptom
	var entitySymptomList []*entity.AdverseSymptom
	for _, symptom := range event.SymptomList {
		entitySymptom := symptom.ToEntity()
		entitySymptom.EventId = entityEvent.Id
		entitySymptomList = append(entitySymptomList, entitySymptom)
	}
	err = a.SymptomRepo.Create(entitySymptomList...)
	if err != nil {
		log.Println(err.Error())
		go func() {
			_ = a.EventRepo.Delete(entityEvent.Id)
		}()
		return errors.New("创建Symptom失败")
	}
	return nil
}

func (a *AdverseEventService) Get(id int64) (*schama.AdverseEvent, error) {
	e, has, err := a.EventRepo.GetById(id)
	if err != nil || !has {
		return nil, errors.New("不良反应报告不存在")
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
	eventList := a.SetUpAdverseEventList(el)
	return eventList, total, nil
}

func (a *AdverseEventService) GetListByUid(uid int64, page, pageSize int) ([]*schama.AdverseEvent, int64, error) {
	el, total, err := a.EventRepo.GetListByUid(uid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	// 获取Vaccine，Symptom
	eventList := a.SetUpAdverseEventList(el)
	return eventList, total, nil
}

func (a *AdverseEventService) CheckAndDelete(id int64, uid int64) error {
	userId, has, err := a.EventRepo.GetUid(id)
	if err != nil || !has {
		return errors.New("不良反应报告不存在")
	}
	if userId != uid {
		return errors.New("不良反应报告不属于您")
	}
	return a.Delete(id)
}

func (a *AdverseEventService) Delete(id int64) error {
	err := a.EventRepo.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除Event失败")
	}
	return nil
}

func (a *AdverseEventService) SetUpAdverseEventList(el []*entity.AdverseEvent) []*schama.AdverseEvent {
	var wg sync.WaitGroup
	var eventList []*schama.AdverseEvent
	for _, e := range el {
		event := &schama.AdverseEvent{}
		event.GetFromEntity(e)
		wg.Add(1)
		go func() {
			_ = a.LoadRequisite(event)
			wg.Done()
		}()
		eventList = append(eventList, event)
	}
	wg.Wait()
	return eventList
}

func (a *AdverseEventService) LoadRequisite(e *schama.AdverseEvent) error {
	// 获取Vaccine，组装VaccineList数据
	vl, err := a.VaccineRepo.GetByEventId(e.Id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("获取Vaccine列表失败")
	}
	var vaccineList []*schama.AdverseVaccine
	for _, v := range vl {
		vaccine := &schama.AdverseVaccine{}
		vax, err := a.VaccineService.Get(v.Id)
		if err != nil {
			log.Println(err.Error())
			return errors.New("获取Vaccine失败")
		}
		vaccine.GetFromEntity(v, vax)
		vaccineList = append(vaccineList, vaccine)
	}
	// 获取Symptom，组装SymptomList数据
	sl, err := a.SymptomRepo.GetByEventId(e.Id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("获取Symptom列表失败")
	}
	var symptomList []*schama.AdverseSymptom
	for _, s := range sl {
		symptom := &schama.AdverseSymptom{}
		symptom.GetFromEntity(s)
		symptomList = append(symptomList, symptom)
	}
	// 组成Event
	e.VaccineList = vaccineList
	e.SymptomList = symptomList
	return nil
}
