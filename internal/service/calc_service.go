package service

import (
	"party-calc/internal/database/repository"
	"sort"
	"time"
)

type PersonData struct {
	Id     int64
	Name   string
	Spent  float64
	Factor int
	Owe    map[string]float64
}

type PersonBalance struct {
	Person *PersonData
	//Person  string
	Balance float64
}

type EventData struct {
	Name    string
	Date    time.Time
	Persons []PersonData
	//Owes map[string]float64
	AllPersonsCount int
	AverageAmount   float64
	TotalAmount     float64
}

type CalcService struct {
	repo     *repository.PersEventsRepository
	service  *PersonsEventsService
	data     *EventData
	balances []PersonBalance
}

func NewCalcService(r *repository.PersEventsRepository, s *PersonsEventsService) *CalcService {
	return &CalcService{repo: r, service: s}
}

func (s *CalcService) createEventData(eventName string) (*EventData, error) {
	ev, _ := s.repo.EventsRepo.Get(&models.Event{Name: eventName})
	data, err := s.service.GetPersonsByEvent(ev.Id)
	if err != nil {
		return err
	}
	s.data = &data
	return nil
}

func (ed *EventData) fillAndSortBalances() {
	for i := 0; i < ed.data.AllPersonsCount-1; i++ {
		ed.balances = append(ed.balances, PersonBalance{
			Person:  &ed.data.Persons[i],
			Balance: ed.data.Persons[i].Spent - ed.data.AverageAmount*float64(ed.data.Persons[i].Factor),
		})
	}
	sort.SliceStable(ed.balances, func(i, j int) bool {
		return ed.balances[i].Balance < ed.balances[j].Balance
	})
}

func (ev *EventData) calculateOwes() {
	for i, j := 0, len(ev.balances)-1; i < j; {
		switch {
		case ev.balances[i].Balance+ev.balances[j].Balance > 0:
			if ev.balances[i].Person.Owe == nil {
				ev.balances[i].Person.Owe = map[string]float64{}
			}
			ev.balances[i].Person.Owe[ev.balances[j].Person.Name] = -ev.balances[i].Balance
			ev.balances[j].Balance += ev.balances[i].Balance
			ev.balances[i].Balance = 0
			i++
		case ev.balances[i].Balance+ev.balances[j].Balance <= 0:
			if ev.balances[i].Person.Owe == nil {
				ev.balances[i].Person.Owe = map[string]float64{}
			}
			ev.balances[i].Person.Owe[ev.balances[j].Person.Name] = ev.balances[j].Balance
			ev.balances[i].Balance += ev.balances[j].Balance
			ev.balances[j].Balance = 0
			j--
		case ev.balances[i].Balance == 0:
			i++
		case ev.balances[j].Balance == 0:
			j--
		}
	}
}


// func (s *CalcService) fillAndSortBalances() {
// 	for i := 0; i < s.data.AllPersonsCount-1; i++ {
// 		s.balances = append(s.balances, PersonBalance{
// 			Person:  &s.data.Persons[i],
// 			Balance: s.data.Persons[i].Spent - s.data.AverageAmount*float64(s.data.Persons[i].Factor),
// 		})
// 	}
// 	sort.SliceStable(s.balances, func(i, j int) bool {
// 		return s.balances[i].Balance < s.balances[j].Balance
// 	})
// }

// func (s *CalcService) calculateOwes() {
// 	for i, j := 0, len(s.balances)-1; i < j; {
// 		switch {
// 		case s.balances[i].Balance+s.balances[j].Balance > 0:
// 			if s.balances[i].Person.Owe == nil {
// 				s.balances[i].Person.Owe = map[string]float64{}
// 			}
// 			s.balances[i].Person.Owe[s.balances[j].Person.Name] = -s.balances[i].Balance
// 			s.balances[j].Balance += s.balances[i].Balance
// 			s.balances[i].Balance = 0
// 			i++
// 		case s.balances[i].Balance+s.balances[j].Balance <= 0:
// 			if s.balances[i].Person.Owe == nil {
// 				s.balances[i].Person.Owe = map[string]float64{}
// 			}
// 			s.balances[i].Person.Owe[s.balances[j].Person.Name] = s.balances[j].Balance
// 			s.balances[i].Balance += s.balances[j].Balance
// 			s.balances[j].Balance = 0
// 			j--
// 		case s.balances[i].Balance == 0:
// 			i++
// 		case s.balances[j].Balance == 0:
// 			j--
// 		}
// 		// if s.balances[i].Balance < s.balances[j].Balance {
// 		// 	s.balances[i].Person.Owe[s.balances[j].Person.Name] = s.balances[i].Balance
// 		// 	s.balances[i].Balance = 0
// 		// 	i++
// 		// } else if s.balances[i].Balance >= s.balances[j].Balance {
// 		// 	s.balances[i].Balance -= s.balances[j].Balance
// 		// 	s.balances[i].Person.Owe[s.balances[j].Person.Name] = s.balances[j].Balance
// 		// 	s.balances[j].Balance = 0
// 		// 	j--
// 		// } else if s.balances[i].Balance == 0 {
// 		// 	i++
// 		// 	continue
// 		// } else if s.balances[j].Balance == 0 {
// 		// 	j--
// 		// 	continue
// 		// }
// 	}
// }

func (s *CalcService) CalcPerson(perName, evName string) (PersonData, error) {
	_, err := s.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return PersonData{}, err
	}
	s.CalcEvent(evName)

	return PersonData{}, nil
}

func (s *CalcService) CalcEvent(name string) (EventData, error) {
	ed,err := s.createEventData(name)
	if err != nil {
		return EventData{}, err
	}
	ed.fillAndSortBalances()
	ed.calculateOwes()
	return *ed, nil
}
