package ambulance

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type implIlnessesAPI struct {
}

// CreateIllness implements IllnessesAPI.
func (i *implIlnessesAPI) CreateIllness(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	patientID := c.Param("patientId")

	patient, err := db.FindDocument(c, patientID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Bad Request",
					"message": "Patient not found",
					"error":   "Wrong patientId probably",
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to retrieve the patient",
					"error":   err.Error(),
				})
		}
		return
	}

	type CreateIllnessRequest struct {
		Diagnosis string `json:"diagnosis" binding:"required"`
		SlFrom    string `json:"sl_from"`
		SlUntil   string `json:"sl_until"`
	}

	var newIllnessRequest CreateIllnessRequest
	if err := c.ShouldBindJSON(&newIllnessRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	illnessID := fmt.Sprintf("I-%d", rand.Intn(90000)+10000)

	newIllness := Illness{
		Id:        illnessID,
		Diagnosis: newIllnessRequest.Diagnosis,
		SlFrom:    newIllnessRequest.SlFrom,
		SlUntil:   newIllnessRequest.SlUntil,
	}

	patient.Illnesses = append(patient.Illnesses, newIllness)

	err = db.UpdateDocument(c, patientID, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to add illness to patient",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, newIllness)
}

// DeleteIllness implements IllnessesAPI.
func (i *implIlnessesAPI) DeleteIllness(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	patientID := c.Param("patientId")

	patient, err := db.FindDocument(c, patientID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Bad Request",
					"message": "Patient not found",
					"error":   "Wrong patientId probably",
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to retrieve the patient",
					"error":   err.Error(),
				})
		}
		return
	}

	illnessId := c.Query("illness_id")

	initialLength := len(patient.Illnesses)
	updatedIllnesses := make([]Illness, 0)

	found := false
	for _, illness := range patient.Illnesses {
		if illness.Id != illnessId {
			updatedIllnesses = append(updatedIllnesses, illness)
		} else {
			found = true
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Illness not found for this patient",
			"error":   "Invalid illness_id",
		})
		return
	}

	if len(updatedIllnesses) == initialLength {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Illness not found",
			"error":   "No illness with the provided ID exists for this patient",
		})
		return
	}

	patient.Illnesses = updatedIllnesses

	err = db.UpdateDocument(c, patientID, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to update patient's illnesses",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Illness deleted successfully",
	})
}

// GetPatientIllnesses implements IllnessesAPI.
func (i *implIlnessesAPI) GetPatientIllnesses(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	patientID := c.Param("patientId")

	patient, err := db.FindDocument(c, patientID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Bad Request",
					"message": "Patient not found",
					"error":   "Wrong patientId probably",
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to retrieve the patient",
					"error":   err.Error(),
				})
		}
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"illnesses": patient.Illnesses,
		},
	)

}

// UpdateSickLeaveEndDate implements IllnessesAPI.
func (i *implIlnessesAPI) UpdateSickLeaveEndDate(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func IllnessesApi() IllnessesAPI {
	return &implIlnessesAPI{}
}
