package models

import (
	"encoding/json"
	"fmt" 
	"log"
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

// UnmarshalJSON personalizado con más depuración
func (m *Match) UnmarshalJSON(data []byte) error {
	// Loguear datos recibidos
	log.Printf("Datos recibidos para deserialización: %s", string(data))

	// Estructura temporal para deserialización
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		log.Printf("Error al deserializar mapa: %v", err)
		return err
	}

	// Extraer valores
	homeTeam, ok := rawData["homeTeam"].(string)
	if !ok {
		log.Printf("homeTeam no es una cadena válida")
		return fmt.Errorf("homeTeam es inválido")
	}
	m.HomeTeam = homeTeam

	awayTeam, ok := rawData["awayTeam"].(string)
	if !ok {
		log.Printf("awayTeam no es una cadena válida")
		return fmt.Errorf("awayTeam es inválido")
	}
	m.AwayTeam = awayTeam

	// Manejar fecha
	matchDateStr, ok := rawData["matchDate"].(string)
	if ok && matchDateStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, matchDateStr)
		if err != nil {
			// Intentar otro formato
			parsedTime, err = time.Parse("2006-01-02", matchDateStr)
			if err != nil {
				log.Printf("Error al parsear fecha: %v", err)
				return err
			}
		}
		m.MatchDate = parsedTime
	} else {
		m.MatchDate = time.Now()
	}

	// Valores opcionales con valores por defecto
	m.Goals = 0
	m.YellowCards = 0
	m.RedCards = 0
	m.ExtraTime = false

	return nil
}