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
	"github.com/gin-gonic/gin"
)

type MedicationsAPI interface {


    // CreateMedication Post /api/patients/:patientId/medications
    // Create a new medication record 
     CreateMedication(c *gin.Context)

    // DeleteMedication Delete /api/patients/:patientId/medications
    // Delete medication record 
     DeleteMedication(c *gin.Context)

    // GetPatientMedications Get /api/patients/:patientId/medications
    // Get medications for a specific patient 
     GetPatientMedications(c *gin.Context)

    // UpdateMedication Patch /api/patients/:patientId/medications/:medicationId
    // Update a medication record 
     UpdateMedication(c *gin.Context)

}