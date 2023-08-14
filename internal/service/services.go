package service

type Services struct {
	PersonService      *PersonService
	EventService       *EventService
	PersonEventService *PersonsEventsService
	CalcService        *CalcService
}

func NewServices(pr PersonRepository, er EventRepository, per PersonsEventsRepository) *Services {
	ps := NewPersonService(pr)
	es := NewEventService(er)
	pes := NewPersonsEventsService(per)
	cs := NewCalcService(ps, es, pes)
	return &Services{
		PersonService:      ps,
		EventService:       es,
		PersonEventService: pes,
		CalcService:        cs,
	}
}
