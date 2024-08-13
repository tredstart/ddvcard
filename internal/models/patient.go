package models

const (
	Papierosy          = 1 << 0
	Alergia            = 1 << 1
	Nadcisnienie       = 1 << 2
	Cukrzyca           = 1 << 3
	ChorobyTarczycy    = 1 << 4
	Epilepsja          = 1 << 5
	SkazaKrowotoczna   = 1 << 6
	Zoltaczka          = 1 << 7
	Gruzlica           = 1 << 8
	AIDS               = 1 << 9
	ChorobyOdogniskowe = 1 << 10
)

type Patient struct {
	Id         string
	Name       string
	Surname    string
	Birthdate  string
	Pesel      string
	Address    string
	Medicines  string
	Illness    int
	Registered string
	Child      int
}

type PatientPartial struct {
	Id        string
	Name      string
	Surname   string
	Child     int
	Pesel     string
	Birthdate string
}
