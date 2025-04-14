package ambulance

import (
	"net/http"

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
	c.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateSickLeaveEndDate implements IllnessesAPI.
func (i *implIlnessesAPI) UpdateSickLeaveEndDate(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func IllnessesApi() IllnessesAPI {
	return &implIlnessesAPI{}
}
