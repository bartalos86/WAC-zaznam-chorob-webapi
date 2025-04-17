package ambulance

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

type implPatientsAPI struct {
}

// CreatePatient implements PatientsAPI.
func (o *implPatientsAPI) CreatePatient(c *gin.Context) {
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

	var requestBody struct {
		Name string `json:"name"`
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

	if requestBody.Name == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Name is required",
				"error":   "name field is empty",
			})
		return
	}

	randomNum := rand.Intn(90000) + 10000
	patientID := fmt.Sprintf("P-%d", randomNum)

	patient := &Patient{
		Name:      requestBody.Name,
		Id:        patientID,
		Illnesses: []Illness{},
	}

	if err := db.CreateDocument(c.Request.Context(), requestBody.Name, patient); err != nil {
		if err == db_service.ErrConflict {
			c.JSON(
				http.StatusConflict,
				gin.H{
					"status":  "Conflict",
					"message": "A patient with this name already exists",
					"error":   err.Error(),
				})
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to create patient",
					"error":   err.Error(),
				})
		}
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  "Created",
			"message": "Patient created successfully",
			"patient": patient,
		})
}

// GetPatients implements PatientsAPI.
func (o *implPatientsAPI) GetPatients(c *gin.Context) {
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

	name := c.Query("name")
	if name != "" {
		patient, err := db.FindByField(c, "name", name)
		if err != nil {
			if err == db_service.ErrNotFound {
				c.JSON(
					http.StatusNotFound,
					gin.H{
						"status":  "Not Found",
						"message": "Patient not found",
						"error":   err.Error(),
					})
				return
			}
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to retrieve patient",
					"error":   err.Error(),
				})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"message":  "Patient retrieved successfully",
			"patients": []Patient{*patient},
		})
		return
	}

	patients, err := db.GetAll(c)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to retrieve patients",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"message":  "Patients retrieved successfully",
		"patients": patients,
	})
}

func (o *implPatientsAPI) DeletePatient(c *gin.Context) {
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

	name := c.Query("name")
	if name != "" {
		patient, err := db.FindByField(c, "name", name)
		if err != nil {
			if err == db_service.ErrNotFound {
				c.JSON(
					http.StatusNotFound,
					gin.H{
						"status":  "Not Found",
						"message": "Patient not found",
						"error":   err.Error(),
					})
				return
			}
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "Internal Server Error",
					"message": "Failed to retrieve patient",
					"error":   err.Error(),
				})
			return
		}

		err = db.DeleteDocument(c, patient.Id)

		if err == nil {
			c.JSON(http.StatusNoContent, gin.H{
				"status":  "No Content",
				"message": "Patient deleted",
			})
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "Failed to delete patient",
				"error":   err.Error(),
			})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "Bad Request",
		"message": "Patient name not provided",
	})
}

func PatientsApi() PatientsAPI {
	return &implPatientsAPI{}
}
