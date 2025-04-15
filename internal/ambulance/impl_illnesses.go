package ambulance

import (
	"net/http"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

type implIlnessesAPI struct {
}

// CreateIllness implements IllnessesAPI.
func (i *implIlnessesAPI) CreateIllness(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

// DeleteIllness implements IllnessesAPI.
func (i *implIlnessesAPI) DeleteIllness(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
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

}

// UpdateSickLeaveEndDate implements IllnessesAPI.
func (i *implIlnessesAPI) UpdateSickLeaveEndDate(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func IllnessesApi() IllnessesAPI {
	return &implIlnessesAPI{}
}
