
package handlers

import (
	"encoding/json" 
	"log"           
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Anton17303/Laboratorio-6/Interno/database"
	"github.com/Anton17303/Laboratorio-6/Interno/Modelo"
)

func GetAllMatches(c *gin.Context) {
	var matches []models.Match
	result := database.GetDB().Find(&matches)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, matches)
}

func GetMatchByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var match models.Match
	result := database.GetDB().First(&match, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.JSON(http.StatusOK, match)
}

func CreateMatch(c *gin.Context) {
	// Capturar el cuerpo raw de la solicitud
	body, err := c.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No se pudo leer el cuerpo de la solicitud",
		})
		return
	}

	// Loguear el contenido raw de la solicitud
	log.Printf("Cuerpo de la solicitud recibido: %s", string(body))

	var newMatch models.Match
	
	// Intentar deserializar manualmente
	if err := json.Unmarshal(body, &newMatch); err != nil {
		log.Printf("Error de deserialización JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error en el formato JSON",
			"details": err.Error(),
		})
		return
	}

	// Validaciones de campos
	if newMatch.HomeTeam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El equipo local es obligatorio"})
		return
	}
	if newMatch.AwayTeam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El equipo visitante es obligatorio"})
		return
	}

	// Si no se proporcionó fecha, usar fecha actual
	if newMatch.MatchDate.IsZero() {
		newMatch.MatchDate = time.Now()
	}

	// Crear el partido
	result := database.GetDB().Create(&newMatch)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newMatch)
}

func UpdateMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	result := database.GetDB().First(&match, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	var updateData models.Match
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar campos específicos
	match.HomeTeam = updateData.HomeTeam
	match.AwayTeam = updateData.AwayTeam
	match.MatchDate = updateData.MatchDate

	result = database.GetDB().Save(&match)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, match)
}

func DeleteMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result := database.GetDB().Delete(&models.Match{}, id)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Partido eliminado exitosamente"})
}

func RegisterGoal(c *gin.Context) {
	updateMatchEvent(c, "goals")
}

func RegisterYellowCard(c *gin.Context) {
	updateMatchEvent(c, "yellow_cards")
}

func RegisterRedCard(c *gin.Context) {
	updateMatchEvent(c, "red_cards")
}

func SetExtraTime(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	
	result := database.GetDB().First(&match, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.ExtraTime = true
	result = database.GetDB().Save(&match)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, match)
}

func updateMatchEvent(c *gin.Context, eventType string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	
	result := database.GetDB().First(&match, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	switch eventType {
	case "goals":
		match.Goals++
	case "yellow_cards":
		match.YellowCards++
	case "red_cards":
		match.RedCards++
	}

	result = database.GetDB().Save(&match)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, match)
}