package person

type Person struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Spent   uint               `json:"spent"`
	Factor  uint               `json:"factor"`  // default must be = 1
	Balance float64            `json:"balance"` // ?? unuseful in output json
	Owe     map[string]float64 `json:"owe"`
}

type Persons struct {
	Persons []Person `json:"persons"`
}

func (per *Person) InitPerson() {
	per.Factor = 1
	per.Owe = map[string]float64{}
}

func (pers *Persons) AddPerson(p Person) {
	pers.Persons = append(pers.Persons, p)
}
