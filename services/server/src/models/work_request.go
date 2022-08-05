package models

import "github.com/google/uuid"

type WorkRequest struct {
	Base
	Hash                 string    `json:"hash" gorm:"uniqueIndex;not null"`
	DifficultyMultiplier int       `json:"difficulty_multiplier"`
	Result               string    `json:"result"`
	Awarded              bool      `json:"awarded" gorm:"default:false;not null"` // Whether or not this has been awarded
	ProvidedBy           uuid.UUID `json:"providedBy" gorm:"not null"`
	RequestedBy          uuid.UUID `json:"requestedBy" gorm:"not null"`
}
