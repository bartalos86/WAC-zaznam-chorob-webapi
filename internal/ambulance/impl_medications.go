/*
 * Patient and Illness Management API
 *
 * Patient and Illness Management API
 *
 * API version: 1.0.0
 * Contact: support@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance

import (
	"net/http"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type implMedicationsAPI struct {
}

// MedicationsApi creates a new instance of MedicationsAPIService
func MedicationsApi() MedicationsAPI {
	return &implMedicationsAPI{}
}

// CreateMedication creates a new medication for a patient
func (i *implMedicationsAPI) CreateMedication(c *gin.Context) {
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

	// Find patient
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

	// Parse medication data
	var medication Medication
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if medication.Name == "" || medication.SideEffects == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Missing required fields",
			"error":   "name and sideEffects are required",
		})
		return
	}

	// Generate a unique ID for the medication
	medicationID := uuid.New().String()
	medication.Id = medicationID

	// Initialize medications array if it's nil
	if patient.Medications == nil {
		patient.Medications = []Medication{}
	}

	// Add the medication to the patient's medications
	patient.Medications = append(patient.Medications, medication)

	// Update patient in database
	err = db.UpdateDocument(c, patientID, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to add medication to patient",
				"error":   err.Error(),
			},
		)
		return
	}

	// Return the created medication
	c.JSON(http.StatusCreated, medication)
}

// GetPatientMedications retrieves all medications for a patient
func (i *implMedicationsAPI) GetPatientMedications(c *gin.Context) {
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

	// Find patient
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

	// Initialize medications array if it's nil
	if patient.Medications == nil {
		patient.Medications = []Medication{}
	}

	// Return the medications
	c.JSON(http.StatusOK, patient.Medications)
}

// UpdateMedication updates an existing medication
func (i *implMedicationsAPI) UpdateMedication(c *gin.Context) {
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
	medicationID := c.Param("medicationId")

	// Find patient
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

	// Parse updated medication data
	var updatedMedication Medication
	if err := c.ShouldBindJSON(&updatedMedication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if updatedMedication.Name == "" || updatedMedication.SideEffects == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Missing required fields",
			"error":   "name and sideEffects are required",
		})
		return
	}

	// Find and update the medication
	found := false
	for i, med := range patient.Medications {
		if med.Id == medicationID {
			updatedMedication.Id = medicationID
			patient.Medications[i] = updatedMedication
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Medication not found",
			"error":   "No medication with the provided ID exists for this patient",
		})
		return
	}

	// Update patient in database
	err = db.UpdateDocument(c, patientID, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to update medication",
				"error":   err.Error(),
			},
		)
		return
	}

	// Return the updated medication
	c.JSON(http.StatusOK, updatedMedication)
}

// DeleteMedication deletes a medication
func (i *implMedicationsAPI) DeleteMedication(c *gin.Context) {
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
	medicationID := c.Query("medication_id")

	if medicationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Missing medication_id query parameter",
			"error":   "medication_id is required",
		})
		return
	}

	// Find patient
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

	// Find and remove the medication
	initialLength := len(patient.Medications)
	updatedMedications := make([]Medication, 0)

	found := false
	for _, med := range patient.Medications {
		if med.Id != medicationID {
			updatedMedications = append(updatedMedications, med)
		} else {
			found = true
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Medication not found for this patient",
			"error":   "Invalid medication_id",
		})
		return
	}

	if len(updatedMedications) == initialLength {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Medication not found",
			"error":   "No medication with the provided ID exists for this patient",
		})
		return
	}

	patient.Medications = updatedMedications

	// Update patient in database
	err = db.UpdateDocument(c, patientID, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to delete medication",
				"error":   err.Error(),
			},
		)
		return
	}

	c.Status(http.StatusNoContent)
}
