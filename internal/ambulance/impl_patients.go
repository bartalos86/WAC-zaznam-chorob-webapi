package ambulance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implPatientsAPI struct {
}

// CreatePatient implements PatientsAPI.
func (o *implPatientsAPI) CreatePatient(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

// GetPatients implements PatientsAPI.
func (o *implPatientsAPI) GetPatients(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func PatientsApi() PatientsAPI {
	return &implPatientsAPI{}
}
