package models

import (
	"ddcard/internal/database"
	"fmt"
	"log"
)

var Illness = map[string]int{
	"Papierosy":          Papierosy,
	"Alergia":            Alergia,
	"Nadcisnienie":       Nadcisnienie,
	"Cukrzyca":           Cukrzyca,
	"ChorobyTarczycy":    ChorobyTarczycy,
	"Epilepsja":          Epilepsja,
	"SkazaKrwotoczna":    SkazaKrwotoczna,
	"Zoltaczka":          Zoltaczka,
	"Gruzlica":           Gruzlica,
	"AIDS":               AIDS,
	"ChorobyOdogniskowe": ChorobyOdogniskowe,
}

const (
	Papierosy          = 1 << 0
	Alergia            = 1 << 1
	Nadcisnienie       = 1 << 2
	Cukrzyca           = 1 << 3
	ChorobyTarczycy    = 1 << 4
	Epilepsja          = 1 << 5
	SkazaKrwotoczna    = 1 << 6
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
	Phone      string
}

type PatientPartial struct {
	Id        string
	Name      string
	Surname   string
	Child     int
	Pesel     string
	Birthdate string
	Phone     string
}

func FetchPatientsByName(name, surname string) ([]PatientPartial, error) {
	var patients []PatientPartial

	if err := database.DB.Select(
		&patients,
		`SELECT id, name, surname, child, pesel, birthdate, phone FROM patient WHERE name LIKE ? OR surname LIKE ?`,
		fmt.Sprintf("%s%%", name),
		fmt.Sprintf("%s%%", surname),
	); err != nil {
		log.Println("couln't read partial by pesel")
		return []PatientPartial{}, err
	}

	return patients, nil
}

func FetchPatientsByPesel(pesel string) ([]PatientPartial, error) {
	var patients []PatientPartial

	if err := database.DB.Select(&patients, `SELECT id, name, surname, child, pesel, birthdate, phone FROM patient WHERE pesel LIKE ?`, fmt.Sprintf("%s%%", pesel)); err != nil {
		log.Println("couln't read partial by pesel")
		return []PatientPartial{}, err
	}

	return patients, nil
}

func PatientCreate(p Patient) error {
	if _, err := database.DB.Exec(
		`INSERT INTO patient VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Id,
		p.Name,
		p.Surname,
		p.Birthdate,
		p.Pesel,
		p.Address,
		p.Medicines,
		p.Illness,
		p.Registered,
		p.Child,
		p.Phone,
	); err != nil {
		return err
	}
	return nil
}
