package storage

import (
	"risks-api/risk"
)

type RiskStorage interface {
	GetAllRisks() []risk.Risk
	// GetRisks(page int, pageSize int) []risk.Risk
	GetRisk(id string) risk.Risk
	InsertRisk(r risk.Risk) error
}
