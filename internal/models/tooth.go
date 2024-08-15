package models

import (
	"ddcard/internal/database"
	"log"
)

const (
	Top int = iota
	Right
	Bottom
	Left
	Center
)

type Tooth struct {
	Id      string
	Patient string
	Number  int
	T       int
	R       int
	B       int
	L       int
	Center  int
}

func FetchTooth(id string) (Tooth, error) {
	var tooth Tooth
	if err := database.DB.Get(&tooth, `SELECT * FROM tooth WHERE id = ?`, id); err != nil {
		log.Println("cannot fetch tooth, ", err)
		return Tooth{}, err
	}
	return tooth, nil
}

func CreateTooth(t Tooth) error {
	if _, err := database.DB.Exec(
		`INSERT INTO tooth VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		t.Id,
		t.Patient,
		t.Number,
		t.T,
		t.R,
		t.B,
		t.L,
		t.Center,
	); err != nil {
		log.Println("cannot create tooth, ", err)
		return err
	}
	return nil
}

func UpdateTooth(t Tooth) error {
	if _, err := database.DB.Exec(
		`UPDATE tooth 
        SET t = ?, r = ?, b = ?, l = ?, center = ?
        WHERE id = ? AND number = ?`,
		t.T, t.R, t.B, t.L, t.Center,
		t.Id, t.Number,
	); err != nil {
		log.Println("cannot update tooth, ", err)
		return err
	}
	return nil
}
