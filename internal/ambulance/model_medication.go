package ambulance

import (
	"encoding/json"
)

// Medication structure corresponds to the Medication schema from the OpenAPI specification
type Medication struct {
	// Unique identifier for the medication record
	Id string `json:"id,omitempty"`
	// Name of the medication
	Name string `json:"name"`
	// Potential side effects of the medication
	SideEffects string `json:"sideEffects"`
}

func (m *Medication) UnmarshalJSON(data []byte) error {
	type Alias Medication
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	return json.Unmarshal(data, aux)
}