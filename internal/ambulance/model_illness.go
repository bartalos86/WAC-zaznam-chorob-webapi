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

type Illness struct {

	// Unique identifier for the illness record
	Id string `json:"id"`

	// Medical diagnosis
	Diagnosis string `json:"diagnosis"`

	// Start date of sick leave
	SlFrom string `json:"sl_from"`

	// End date of sick leave
	SlUntil string `json:"sl_until"`

	// List of treatments related to this illness
	Treatments []Treatment `json:"treatments"`
}
