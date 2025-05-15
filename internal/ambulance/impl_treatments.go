package ambulance

import (
	"net/http"
	"slices"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type implTreatmentsAPI struct {
}

// CreateTreatment implements TreatmentsAPI.
func (o *implTreatmentsAPI) CreateTreatment(c *gin.Context) {
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

	patientId := c.Param("patientId")
	illnessId := c.Param("illnessId")

	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   "Invalid patientId",
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

	var requestBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	treatmentId := uuid.New().String()

	newTreatment := &Treatment{
		Id:          treatmentId,
		Name:        requestBody.Name,
		Description: requestBody.Description,
		StartDate:   requestBody.StartDate,
		EndDate:     requestBody.EndDate,
	}

	found := false
	for idx := range patient.Illnesses {
		if patient.Illnesses[idx].Id == illnessId {
			patient.Illnesses[idx].Treatments = append(patient.Illnesses[idx].Treatments, *newTreatment)
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Illness not found for this patient",
			"error":   "Invalid illnessId",
		})
		return
	}

	err = db.UpdateDocument(c, patientId, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to add treatment to patient's illness",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":    "Created",
			"message":   "Treatment created successfully",
			"treatment": newTreatment,
		})
}

// GetTreatments implements TreatmentsAPI.
func (o *implTreatmentsAPI) GetTreatments(c *gin.Context) {
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

	patientId := c.Param("patientId")
	illnessId := c.Param("illnessId")

	patient, err := db.FindByField(c, "id", patientId)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			})
		return
	}

	var illness *Illness
	for _, ill := range patient.Illnesses {
		if ill.Id == illnessId {
			illness = &ill
			break
		}
	}

	if illness == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Illness not found for this patient",
				"error":   "Illness not found",
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"message":    "Treatments retrieved successfully",
		"treatments": illness.Treatments,
	})
}

// UpdateTreatment implements TreatmentsAPI.
func (o *implTreatmentsAPI) UpdateTreatment(c *gin.Context) {
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

	patientId := c.Param("patientId")
	illnessId := c.Param("illnessId")
	treatmentId := c.Param("treatmentId")

	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   "Invalid patientId",
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

	var requestBody struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	foundIllness := false
	foundTreatment := false
	for idx := range patient.Illnesses {
		if patient.Illnesses[idx].Id == illnessId {
			foundIllness = true
			for idy := range patient.Illnesses[idx].Treatments {
				if patient.Illnesses[idx].Treatments[idy].Id == treatmentId {
					foundTreatment = true
					patient.Illnesses[idx].Treatments[idy].Name = requestBody.Name
					patient.Illnesses[idx].Treatments[idy].Description = requestBody.Description
					patient.Illnesses[idx].Treatments[idy].StartDate = requestBody.StartDate
					patient.Illnesses[idx].Treatments[idy].EndDate = requestBody.EndDate
					break
				}
			}
			break
		}
	}

	if !foundIllness {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Illness not found for this patient",
			"error":   "Invalid illnessId",
		})
		return
	}
	if !foundTreatment {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Treatment for illness not found for this patient",
			"error":   "Invalid treatmentId",
		})
		return
	}

	err = db.UpdateDocument(c, patientId, patient)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to update treatment",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Treatment updated successfully",
	})
}

// DeleteTreatment implements TreatmentsAPI.
func (o *implTreatmentsAPI) DeleteTreatment(c *gin.Context) {
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

	patientId := c.Param("patientId")
	illnessId := c.Param("illnessId")
	treatmentId := c.Param("treatmentId")

	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   "Invalid patientId",
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

	illnessIndex := -1
	treatmentIndex := -1
	for idx := range patient.Illnesses {
		if patient.Illnesses[idx].Id == illnessId {
			illnessIndex = idx
			for idy := range patient.Illnesses[idx].Treatments {
				if patient.Illnesses[idx].Treatments[idy].Id == treatmentId {
					treatmentIndex = idy
					break
				}
			}
			break
		}
	}

	if illnessIndex == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Illness not found for this patient",
			"error":   "Invalid illnessId",
		})
		return
	}
	if treatmentIndex == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Treatment for illness not found for this patient",
			"error":   "Invalid treatmentId",
		})
		return
	} else {
		// Remove treatment from the array
		patient.Illnesses[illnessIndex].Treatments = slices.Delete(
			patient.Illnesses[illnessIndex].Treatments, treatmentIndex,
			treatmentIndex+1,
		)
	}

	err = db.UpdateDocument(c, patientId, patient)
	if err == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status":  "No Content",
			"message": "Treatment deleted",
		})
		return
	}

	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to delete treatment",
			"error":   err.Error(),
		})
}

// TreatmentsApi returns an implementation of TreatmentsAPI.
func TreatmentsApi() TreatmentsAPI {
	return &implTreatmentsAPI{}
}
