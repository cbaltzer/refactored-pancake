package storage

import (
	"fmt"
	"log"
	"risks-api/risk"
)

func NewInMemoryStorage() InMemory {
	return InMemory{
		risks:   []risk.Risk{},
		indexes: map[string]int{},
	}
}

func (db *InMemory) GetAllRisks() []risk.Risk {
	return db.risks
}

// func (db InMemory) GetRisks(page int, pageSize int) []risk.Risk {
// 	return []risk.Risk{}
// }

func (db *InMemory) GetRisk(id string) (*risk.Risk, error) {

	if idx, ok := db.indexes[id]; ok {
		if idx < len(db.risks) {
			risk := db.risks[idx]

			log.Printf("Found risk ID %s: %v", id, risk)
			return &risk, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (db *InMemory) PutRisk(r risk.Risk) error {
	if _, ok := db.indexes[r.ID]; ok {
		return fmt.Errorf("already exists")
	}

	idx := len(db.risks)
	db.risks = append(db.risks, r)
	db.indexes[r.ID] = idx

	log.Printf("Inserted risk ID %s [internal index: %d]", r.ID, idx)
	return nil
}
