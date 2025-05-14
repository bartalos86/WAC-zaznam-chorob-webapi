package seeder

import (
	"context"
	"math/rand"

	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/ambulance"
	"github.com/bartalos86/WAC-zaznam-chorob-webapi/internal/db_service"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func Seed(db db_service.DbService[ambulance.Patient]) {
	ctx := context.Background()

	for range 15 {
		numIllnesses := rand.Intn(6) // 0–5

		var illnesses []ambulance.Illness
		for j := 0; j < numIllnesses; j++ {

			numTreatments := rand.Intn(4) // 0-3
			var treatments []ambulance.Treatment
			for k := 0; k < numTreatments; k++ {

				treatment := ambulance.Treatment{
					Id:          uuid.New().String(),
					Name:        faker.Word(),
					Description: faker.Sentence(),
					StartDate:   faker.Date(),
					EndDate:     faker.Date(),
				}
				treatments = append(treatments, treatment)
			}

			illness := ambulance.Illness{
				Id:         uuid.New().String(),
				Diagnosis:  faker.Word(),
				SlFrom:     faker.Date(),
				SlUntil:    faker.Date(),
				Treatments: treatments,
			}
			illnesses = append(illnesses, illness)
		}

		// Generate random medications for the patient
		numMedications := rand.Intn(4) // 0-3
		var medications []ambulance.Medication
		for j := 0; j < numMedications; j++ {
			medication := ambulance.Medication{
				Id:          uuid.New().String(),
				Name:        faker.Word(),
				SideEffects: faker.Sentence(),
			}
			medications = append(medications, medication)
		}

		patient := ambulance.Patient{
			Id:          uuid.New().String(),
			Name:        faker.Name(),
			Illnesses:   illnesses,
			Medications: medications,
		}

		db.CreateDocument(ctx, patient.Name, &patient)
	}
}
