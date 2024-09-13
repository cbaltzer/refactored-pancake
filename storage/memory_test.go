package storage_test

import (
	"fmt"
	"risks-api/risk"
	"risks-api/storage"
	"testing"
)

var db = storage.NewInMemoryStorage()

func TestInsertRisk(t *testing.T) {
	r := risk.NewRisk("test", "test desc")

	if len(db.GetAllRisks()) != 0 {
		t.Errorf("db should initialize empty")
	}

	db.InsertRisk(r)

	if len(db.GetAllRisks()) != 1 {
		t.Errorf("db insert failed")
	}
}

func TestGetRisk(t *testing.T) {
	r := risk.NewRisk("test", "test desc")
	db.InsertRisk(r)

	get, err := db.GetRisk(r.ID)
	if err != nil {
		t.Errorf("db get failed: %s", err)
	}

	if *get != r {
		t.Errorf("db get didn't return match")
	}
}

func TestGetAllRisks(t *testing.T) {
	var db = storage.NewInMemoryStorage()

	r1 := risk.NewRisk("test1", "test desc")
	r2 := risk.NewRisk("test2", "test desc")
	r3 := risk.NewRisk("test3", "test desc")

	mockRes := []risk.Risk{r1, r2, r3}

	db.InsertRisk(r1)
	db.InsertRisk(r2)
	db.InsertRisk(r3)

	res := db.GetAllRisks()
	if len(res) != 3 {
		t.Errorf("db didn't return correct size")
	}

	fmt.Printf("%v\n", mockRes)
	fmt.Printf("%v\n", res)

	for i := range res {
		if mockRes[i] != res[i] {
			t.Errorf("in/output mismatch")
		}
	}
}
