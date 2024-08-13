package routes_test

import (
	"ddcard/internal/models"
	"ddcard/internal/server/routes"
	"testing"
)

func assert(exp bool, msg string, t *testing.T) {
	if !exp {
		t.Error(msg)
	}
}

func TestHappyPathOneByOne(t *testing.T) {
	assert(routes.GetIllness([]string{"Papierosy"}) == models.Papierosy, "papierosy not worky", t)
	assert(routes.GetIllness([]string{"Alergia"}) == models.Alergia, "alergia not worky", t)
	assert(routes.GetIllness([]string{"Nadcisnienie"}) == models.Nadcisnienie, "nadcisnienie not worky", t)
	assert(routes.GetIllness([]string{"Cukrzyca"}) == models.Cukrzyca, "cukrzyca not worky", t)
	assert(routes.GetIllness([]string{"ChorobyTarczycy"}) == models.ChorobyTarczycy, "ChorobyTarczycy not worky", t)
	assert(routes.GetIllness([]string{"Epilepsja"}) == models.Epilepsja, "Epilepsja not worky", t)
	assert(routes.GetIllness([]string{"SkazaKrwotoczna"}) == models.SkazaKrwotoczna, "SkazaKrwotoczna not worky", t)
	assert(routes.GetIllness([]string{"Zoltaczka"}) == models.Zoltaczka, "Zoltaczka not worky", t)
	assert(routes.GetIllness([]string{"Gruzlica"}) == models.Gruzlica, "Gruzlica not worky", t)
	assert(routes.GetIllness([]string{"AIDS"}) == models.AIDS, "AIDS not worky", t)
	assert(routes.GetIllness([]string{"ChorobyOdogniskowe"}) == models.ChorobyOdogniskowe, "ChorobyOdogniskowe not worky", t)
}

func TestMulti(t *testing.T) {
	case0 := []string{}
	case1 := []string{"Papierosy", "Alergia"}
	case2 := []string{"Nadcisnienie", "Cukrzyca", "ChorobyTarczycy", "AIDS"}
	case3 := []string{"Epilepsja", "SkazaKrwotoczna", "Zoltaczka", "Gruzlica", "ChorobyOdogniskowe"}

	assert(routes.GetIllness(case0) == 0, "empty case failed", t)
	assert(routes.GetIllness(case1) == models.Papierosy|models.Alergia, "case 1 failed", t)
	assert(routes.GetIllness(case2) == models.Nadcisnienie|models.Cukrzyca|models.ChorobyTarczycy|models.AIDS, "case 2 failed", t)
	assert(routes.GetIllness(case3) == models.Epilepsja|models.SkazaKrwotoczna|models.Zoltaczka|models.Gruzlica|models.ChorobyOdogniskowe, "case 3 failed", t)

	case_wrong := []string{"wrong case"}
	assert(routes.GetIllness(case_wrong) == 0, "wrong case failed", t)
	case_wrong_in_front := []string{"wrong case", "Papierosy"}
	assert(routes.GetIllness(case_wrong_in_front) == models.Papierosy, "case wrong in front", t)
	case_wrong_back := []string{"Papierosy", "wrong case"}
	assert(routes.GetIllness(case_wrong_back) == models.Papierosy, "case wrong back", t)
	case_wrong_mid := []string{"Papierosy", "wrong case", "Alergia"}
	assert(routes.GetIllness(case_wrong_mid) == models.Papierosy|models.Alergia, "mid case failed ", t)
}
