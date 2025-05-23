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

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the IllnessesAPI part of the API
	IllnessesAPI IllnessesAPI
	// Routes for the MedicationsAPI part of the API
	MedicationsAPI MedicationsAPI
	// Routes for the PatientsAPI part of the API
	PatientsAPI PatientsAPI
	// Routes for the TreatmentsAPI part of the API
	TreatmentsAPI TreatmentsAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"CreateIllness",
			http.MethodPost,
			"/api/patients/:patientId/illnesses",
			handleFunctions.IllnessesAPI.CreateIllness,
		},
		{
			"DeleteIllness",
			http.MethodDelete,
			"/api/patients/:patientId/illnesses",
			handleFunctions.IllnessesAPI.DeleteIllness,
		},
		{
			"GetPatientIllnesses",
			http.MethodGet,
			"/api/patients/:patientId/illnesses",
			handleFunctions.IllnessesAPI.GetPatientIllnesses,
		},
		{
			"UpdateSickLeaveEndDate",
			http.MethodPatch,
			"/api/patients/:patientId/illnesses",
			handleFunctions.IllnessesAPI.UpdateSickLeaveEndDate,
		},
		{
			"CreateMedication",
			http.MethodPost,
			"/api/patients/:patientId/medications",
			handleFunctions.MedicationsAPI.CreateMedication,
		},
		{
			"DeleteMedication",
			http.MethodDelete,
			"/api/patients/:patientId/medications",
			handleFunctions.MedicationsAPI.DeleteMedication,
		},
		{
			"GetPatientMedications",
			http.MethodGet,
			"/api/patients/:patientId/medications",
			handleFunctions.MedicationsAPI.GetPatientMedications,
		},
		{
			"UpdateMedication",
			http.MethodPatch,
			"/api/patients/:patientId/medications/:medicationId",
			handleFunctions.MedicationsAPI.UpdateMedication,
		},
		{
			"CreatePatient",
			http.MethodPost,
			"/api/patients",
			handleFunctions.PatientsAPI.CreatePatient,
		},
		{
			"DeletePatient",
			http.MethodDelete,
			"/api/patients",
			handleFunctions.PatientsAPI.DeletePatient,
		},
		{
			"GetPatients",
			http.MethodGet,
			"/api/patients",
			handleFunctions.PatientsAPI.GetPatients,
		},
		{
			"CreateTreatment",
			http.MethodPost,
			"/api/patients/:patientId/illnesses/:illnessId/treatments",
			handleFunctions.TreatmentsAPI.CreateTreatment,
		},
		{
			"DeleteTreatment",
			http.MethodDelete,
			"/api/patients/:patientId/illnesses/:illnessId/treatments/:treatmentId",
			handleFunctions.TreatmentsAPI.DeleteTreatment,
		},
		{
			"GetTreatments",
			http.MethodGet,
			"/api/patients/:patientId/illnesses/:illnessId/treatments",
			handleFunctions.TreatmentsAPI.GetTreatments,
		},
		{
			"UpdateTreatment",
			http.MethodPatch,
			"/api/patients/:patientId/illnesses/:illnessId/treatments/:treatmentId",
			handleFunctions.TreatmentsAPI.UpdateTreatment,
		},
	}
}
