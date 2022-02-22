package location

import "fmt"

var (
	ErrCepSyntax   = fmt.Errorf("CEP syntax error")
	ErrCepNotFound = fmt.Errorf("CEP not found")
	ErrCepEmpty    = fmt.Errorf("CEP empty")
)

type Location struct {
	CEP       string `json:"cep"`
	CepStatus bool   `json:"cep_status"`
	StatusMsg string `json:"status_msg"`
	Street    string `json:"street"`
	Number    []int  `json:"number"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
}

func (l *Location) CheckCEP() {

	l.CepStatus = true
	l.StatusMsg = "status: ok"

	if len(l.CEP) != 8 {
		l.CepStatus = false
		l.StatusMsg = "must be 8 characters"
	}
}

func (l *Location) SetCEP(c string) {

	l.CEP = c
	l.CheckCEP()
}

func (l *Location) GetCEP() string {
	return l.CEP
}

func NewLocation() Location {
	l := Location{}
	l.CepStatus = false
	return l
}
