package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	HomeTeam    string    `json:"homeTeam" gorm:"not null"`
	AwayTeam    string    `json:"awayTeam" gorm:"not null"`
	MatchDate   time.Time `json:"matchDate" gorm:"not null"`
	Goals       int       `json:"goals" gorm:"default:0"`
	YellowCards int       `json:"yellowCards" gorm:"default:0"`
	RedCards    int       `json:"redCards" gorm:"default:0"`
	ExtraTime   bool      `json:"extraTime" gorm:"default:false"`
}