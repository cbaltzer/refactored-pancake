package storage

import (
	"risks-api/risk"
)

type RiskStorage interface {
	GetAllRisks() []risk.Risk
	// GetRisks(page int, pageSize int) []risk.Risk
	GetRisk(id string) risk.Risk
	PutRisk(r risk.Risk) error
}

type InMemory struct {
	risks   []risk.Risk
	indexes map[string]int
}
