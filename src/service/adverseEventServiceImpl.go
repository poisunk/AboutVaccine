package service

import (
	"about-vaccine/src/models"
	"about-vaccine/src/utile"
	"errors"
	"log"
	"sync"
	"time"
)

type AdverseEventServiceImpl struct {
	VaccineService
}

func InitAdverseEventService() AdverseEventService {
	return &AdverseEventServiceImpl{
		VaccineService: InitVaccineService(),
	}
}

func (service *AdverseEventServiceImpl) CreateAdverseEvent(event AdverseEvent) error {
	// 准备数据，AdverseEvent
	adverseEvent := &models.AdverseEvent{}
	err := utile.StructConv(event, adverseEvent)
	if err != nil {
		log.Println(err.Error())
		return errors.New("数据格式有误！")
	}
	adverseEvent.CreateDate = time.Now()
	log.Println(adverseEvent.CreateDate)
	// 提交数据
	err = models.CreateAdverseEvent(adverseEvent)
	// 提交失败
	if err != nil {
		log.Println(err.Error())
		return errors.New("提交失败！")
	}

	// 准备数据，AdverseVaccine
	var adverseVaccineList []*models.AdverseVaccine
	for _, v := range event.VaccineList {
		av := &models.AdverseVaccine{
			AdverseEventId: adverseEvent.Id,
			VaccineId:      v.Id,
			VaccinateDate:  v.VaccinateDate,
			Dose:           v.Dose,
			Route:          v.Route,
			Site:           v.Site,
		}
		adverseVaccineList = append(adverseVaccineList, av)
	}
	// 提交数据
	err = models.CreateAdverseVaccineList(adverseVaccineList)
	if err != nil {
		log.Println(err.Error())
		// 如果提交失败，则把之前的记录也一并删除
		go func() {
			_ = models.DeleteAdverseEventById(adverseEvent.Id)
		}()
		return errors.New("创建失败！")
	}

	// 准备数据，AdverseSymptom
	var adverseSymptomList []*models.AdverseSymptom
	for _, v := range event.SymptomList {
		as := &models.AdverseSymptom{
			EventId: adverseEvent.Id,
			Symptom: v.Symptom,
			OaeId:   v.OaeId,
		}
		adverseSymptomList = append(adverseSymptomList, as)
	}
	// 提交数据
	err = models.CreateAdverseSymptomList(adverseSymptomList)
	if err != nil {
		log.Println(err.Error())
		// 如果提交失败，则把之前的记录也一并删除
		go func() {
			_ = models.DeleteAdverseEventById(adverseEvent.Id)
		}()
		return errors.New("创建失败！")
	}
	return nil
}

func (service *AdverseEventServiceImpl) GetAdverseEvent(id int64) (event *AdverseEvent, err error) {
	// 检索AdverseEvent
	adverseEvent, err := models.GetAdverseEventById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("检索失败！")
	}

	// 检索AdverseVaccine
	adverseVaccineList, err := models.GetAdverseVaccineListByVid(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("检索失败！")
	}

	// 检索AdverseSymptom
	adverseSymptomList, err := models.GetAdverseSymptomByEventId(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("检索失败！")
	}

	// 组装数据
	// 组装AdverseEvent部分
	event = &AdverseEvent{}
	err = utile.StructConv(adverseEvent, event)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("数据格式有误！")
	}
	// 组装AdverseVaccine部分
	var vaccineList []*AdverseVaccine
	wg := &sync.WaitGroup{}
	wg.Add(len(adverseVaccineList))
	for _, v := range adverseVaccineList {
		go func(v *models.AdverseVaccine) {
			defer wg.Done()
			vaccine, err := service.GetVaccine(v.VaccineId)
			if err != nil {
				log.Println(err.Error())
				return
			}
			vaccineList = append(vaccineList, &AdverseVaccine{
				Id:           v.VaccineId,
				Type:         vaccine.Type,
				Manufacturer: vaccine.ProductionCompany,
				Name:         vaccine.ProductName,
				Dose:         v.Dose,
				Route:        v.Route,
				Site:         v.Site,
			})
		}(v)
	}
	wg.Wait()
	event.VaccineList = vaccineList

	// 组装AdverseSymptom部分
	var symptomList []*AdverseSymptom
	for _, v := range adverseSymptomList {
		symptom := &AdverseSymptom{
			Id:      v.Id,
			EventId: v.EventId,
			Symptom: v.Symptom,
			OaeId:   v.OaeId,
		}
		symptomList = append(symptomList, symptom)
	}
	event.SymptomList = symptomList
	return event, nil
}

func (service *AdverseEventServiceImpl) DeleteAdverseEvent(id int64) error {
	// 删除AdverseEvent
	err := models.DeleteAdverseEventById(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除失败！")
	}
	return nil
}

func (service *AdverseEventServiceImpl) GetAdverseEventList(page, pageSize int) (
	adverseEventList []*AdverseEvent, total int64, err error) {
	// 检索AdverseEvent
	list, err := models.GetAdverseEventList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("检索失败！")
	}
	// 组装数据
	for _, v := range list {
		event, err := service.GetAdverseEvent(v.Id)
		if err != nil {
			return nil, 0, err
		}
		adverseEventList = append(adverseEventList, event)
	}
	total, err = models.CountAdverseEvent()
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("检索失败！")
	}
	return adverseEventList, total, nil
}
