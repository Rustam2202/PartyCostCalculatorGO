package service

import (
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
	"sort"
	"time"
)

type PersonData struct {
	Id    int64
	Name  string
	Spent float64
	Owe   map[string]float64
}

type PersonBalance struct {
	Person  *PersonData
	Balance float64
}

type EventData struct {
	Name            string
	Date            time.Time
	Owes            []PersonData
	AllPersonsCount int
	AverageAmount   float64
	TotalAmount     float64
}

type CalcService struct {
	repo    *repository.PersEventsRepository
	service *PersEventsService
	//data    *EventData
}

func NewCalcService(r *repository.PersEventsRepository, s *PersEventsService) *CalcService {
	return &CalcService{repo: r, service: s}
}


func (s *CalcService) CreateEventData(eventName string) (EventData, error) {
	ev, _ := s.repo.EventsRepo.Get(&models.Event{Name: eventName})
	result, err := s.service.GetPersonsByEvent(ev.Id)
	if err != nil {
		return EventData{}, err
	}
	return result, nil
}

func fillAndSortBalances(data *EventData) []PersonBalance {
	pers := data.Owes
	var result []PersonBalance
	for _, p := range pers {
		result = append(result, PersonBalance{
			Person:  &p,
			Balance: p.Spent,
		})
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Balance < result[j].Balance
	})
	return result
}

func (s *CalcService) calculateOwes(b []PersonBalance) {
	for i, j := 0, len(b)-1; i < j; {
		if b[i].Balance < b[j].Balance {
			b[i].Person.Owe[b[j].Person.Name] = b[i].Balance
			b[i].Balance = 0
			i++
		} else if b[i].Balance >= b[j].Balance {
			b[i].Balance -= b[j].Balance
			b[i].Person.Owe[b[j].Person.Name] = b[j].Balance
			b[j].Balance = 0
			j--
		} else if b[i].Balance == 0 {
			i++
			continue
		} else if b[j].Balance == 0 {
			j--
			continue
		}
	}
}

func (s *CalcService) CalcPerson(name string) (PersonData, error) {
	per, err := s.repo.PersRepo.Get(&models.Person{Name: name})
	if err!=nil{
		return PersonData{},err
	}
	perEv, err := s.repo.GetPerson(&models.PersonsAndEvents{PersonId: per.Id})
	if err!=nil{
		return PersonData{},err
	}
	ev, err := s.repo.EventsRepo.Get(&models.Event{Id: perEv.EventId})
	if err!=nil{
		return PersonData{},err
	}
	s.CalcEvent(ev.Name)
	
	return PersonData{}, nil
}

func (s *CalcService) CalcEvent(name string) ( error) {
	data, _ := s.CreateEventData(name)
	balances := fillAndSortBalances(&data)
	s.calculateOwes(balances)
	return nil
}
