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

	serviceD := factory.NewServiceFactory().
		WithName("Threat Intelligence Feed").
		WithCategory("Security").
		WithCharacteristic("Update Frequency", "Every 1h").
		Build()
	_, _ = collection.InsertOne(ctx, serviceD)

	serviceC := factory.NewServiceFactory().
		WithName("Firewall Protection").
		WithCategory("Security").
		WithCharacteristic("Policy Level", "High").
		WithRelationshipTo(*serviceD.ID, "Threat Intelligence Feed", "reliesOn").
		Build()
	_, _ = collection.InsertOne(ctx, serviceC)

	serviceB := factory.
		NewServiceFactory().
		WithName("Router Configuration").
		WithCategory("Infrastructure").
		WithCharacteristic("Router Model", "Cisco XR500").
		Build()
	_, _ = collection.InsertOne(ctx, serviceB)

	// A zavisi od B
	serviceA := factory.NewServiceFactory().
		WithName("Internet Access Service").
		WithCategory("Business").
		WithCharacteristic("bandwidth", fmt.Sprintf("%d", gofakeit.Number(50, 1000))).
		WithCharacteristic("unit", "Mbps").
		WithCharacteristic("technology", "Fiber").
		WithRelationshipTo(*serviceB.ID, "Router Configuration", "reliesOn").
		WithRelationshipTo(*serviceC.ID, "Firewall Protection", "bundledWith").
		Build()

	docs := []interface{}{serviceB, serviceA}

	_, err := collection.InsertMany(ctx, docs)
	return err
}
