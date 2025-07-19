package seeders

import (
	"context"
	"fmt"
	factory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/database/factories"
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedServices(collection *mongo.Collection, count int) error {
	ctx := context.TODO()
	category := []string{"Network", "Infrastructure", "BSOD", "XGS-PON", "I-VPN"}
	services := factory.CreateMany(count, func(i int) *factory.ServiceFactory {
		return factory.NewServiceFactory().
			WithCategory(gofakeit.RandomString(category)).
			WithCharacteristic("bandwidth", fmt.Sprintf("%d Mbps", gofakeit.Number(50, 1000))).
			WithRelatedParty("Owner", gofakeit.Name())
	})

	docs := make([]interface{}, len(services))
	for i, s := range services {
		docs[i] = s
	}

	_, err := collection.InsertMany(ctx, docs)
	return err
}

func SeedServicesWithRelationshipTo(collection *mongo.Collection, deep int) error {
	ctx := context.TODO()

	// Prvo kreiramo B (zavisni servis)
	serviceB := factory.NewServiceFactory().WithCategory("Infrastructure").Build()

	// A zavisi od B
	serviceA := factory.NewServiceFactory().
		WithCategory("Network").
		WithCharacteristic("bandwidth", fmt.Sprintf("%d", gofakeit.Number(50, 1000))).
		WithCharacteristic("unit", "Mbps").
		WithRelationshipTo(*serviceB.ID, "dependsOn").
		Build()

	docs := []interface{}{serviceB, serviceA}

	_, err := collection.InsertMany(ctx, docs)
	return err
}
