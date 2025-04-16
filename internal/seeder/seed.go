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

	for i := 0; i < 15; i++ {
		numIllnesses := rand.Intn(6) // 0–5

		var illnesses []ambulance.Illness
		for j := 0; j < numIllnesses; j++ {
			illness := ambulance.Illness{
				Id:        uuid.New().String(),
				Diagnosis: faker.Word(),
				SlFrom:    faker.Date(),
				SlUntil:   faker.Date(),
			}
			illnesses = append(illnesses, illness)
		}

		patient := ambulance.Patient{
			Id:        uuid.New().String(),
			Name:      faker.Name(),
			Illnesses: illnesses,
		}

		db.CreateDocument(ctx, patient.Name, &patient)
	}
}
