package risk

import (
	"github.com/google/uuid"
)

type Risk struct {
	ID          string     `json:"id"`
	Status      RiskStatus `json:"status"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}

type RiskStatus string

const (
	RiskStatusOpen          RiskStatus = "open"
	RiskStatusClosed        RiskStatus = "closed"
	RiskStatusAccepted      RiskStatus = "accepted"
	RiskStatusInvestigating RiskStatus = "investigating"
)

func NewRisk(title string, description string) Risk {
	r := Risk{
		ID:          uuid.New().String(),
		Status:      RiskStatusOpen,
		Title:       title,
		Description: description,
	}

	return r
}
