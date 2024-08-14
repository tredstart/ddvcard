package models

import (
	"ddcard/internal/database"
	"log"
)

type Tooth struct {
	Id      string
	Patient string
	Number  int
	Top     int
	Right   int
	Bottom  int
	Left    int
	Center  int
}

func CreateTooth(t Tooth) error {
	if _, err := database.DB.Exec(
		`INSERT INTO tooth VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		t.Id,
		t.Patient,
		t.Number,
		t.Top,
		t.Right,
		t.Bottom,
		t.Left,
		t.Center,
	); err != nil {
		log.Println("cannot create tooth")
		return err
	}
	return nil
}
