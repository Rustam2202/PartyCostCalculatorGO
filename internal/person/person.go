package person

type Person struct {
	id           int
	Name         string `json:"name"`
	Spent        uint   `json:"spent"`
	Participants uint   `json:"participants"`
	Balance      float32
	IndeptedTo   map[string]float32
}

type Persons struct {
	Persons []Person `json:"persons"`
}
