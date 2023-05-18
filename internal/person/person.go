package person

type Person struct {
	Name       string `json:"name"`
	Spent      uint   `json:"spent"`
	Factor     uint   `json:"factor"` // default must be = 1
	Balance    float64
	IndeptedTo map[string]float64
}

type Persons struct {
	Persons []Person `json:"persons"`
}

func (per *Person) InitPerson() {
	per.Factor = 1
	per.IndeptedTo = map[string]float64{}
}

func (pers *Persons) AddPerson(p Person) {
	pers.Persons = append(pers.Persons, p)
}
