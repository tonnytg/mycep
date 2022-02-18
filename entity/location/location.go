package location

type Location struct {
	CEP string `json:"cep"`
	Street string `json:"street"`
	Number []int `json:"number"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}

func (l *Location) SetCEP(c string) error {
	l.CEP = c
	return nil
}

func (l *Location) GetCEP() string {
	return l.CEP
}

func NewLocation() Location {
	return Location{}
}
