package seeders

import (
	"context"
	"fmt"
	factory "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/database/factories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
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

//func SeedServicesWithRelationshipTo(collection *mongo.Collection, count int, deep int, rootName string) error {
//	ctx := context.TODO()
//
//	serviceD := factory.NewServiceFactory().
//		WithName("Threat Intelligence Feed").
//		WithCategory("Security").
//		WithCharacteristic("Update Frequency", "Every 1h").
//		Build()
//	_, _ = collection.InsertOne(ctx, serviceD)
//
//	serviceC := factory.NewServiceFactory().
//		WithName("Firewall Protection").
//		WithCategory("Security").
//		WithCharacteristic("Policy Level", "High").
//		WithRelationshipTo(*serviceD.ID, "Threat Intelligence Feed", "reliesOn").
//		Build()
//	_, _ = collection.InsertOne(ctx, serviceC)
//
//	serviceB := factory.
//		NewServiceFactory().
//		WithName("Router Configuration").
//		WithCategory("Infrastructure").
//		WithCharacteristic("Router Model", "Cisco XR500").
//		Build()
//	_, _ = collection.InsertOne(ctx, serviceB)
//
//	// A zavisi od B
//	serviceA := factory.NewServiceFactory().
//		WithName("Internet Access Service").
//		WithCategory("Business").
//		WithCharacteristic("bandwidth", fmt.Sprintf("%d", gofakeit.Number(50, 1000))).
//		WithCharacteristic("unit", "Mbps").
//		WithCharacteristic("technology", "Fiber").
//		WithRelationshipTo(*serviceB.ID, "Router Configuration", "reliesOn").
//		WithRelationshipTo(*serviceC.ID, "Firewall Protection", "bundledWith").
//		Build()
//
//	docs := []interface{}{serviceA}
//
//	_, err := collection.InsertMany(ctx, docs)
//	return err
//}

func SeedServicesWithRelationshipTo(collection *mongo.Collection, count int, deep int, rootName string) error {
	ctx := context.TODO()

	if count <= 0 {
		count = 1
	}
	if deep <= 0 {
		deep = 2
	}

	var docs []interface{}

	for i := 0; i < count; i++ {
		name := gofakeit.ProductName()
		if rootName != "" && i == 0 {
			name = rootName
		}

		levelServices := []*models.Service{}
		var count = 0
		for d := deep; d > 0; d-- {

			var child *models.Service
			if d == deep {
				child = factory.NewServiceFactory().
					WithName(fmt.Sprintf("Service Child %d-%d", i+1, d)).
					WithCategory(gofakeit.RandomString([]string{"Security", "Infrastructure", "Monitoring"})).
					WithCharacteristic("version", gofakeit.AppVersion()).
					Build()

			} else {
				child = factory.NewServiceFactory().
					WithName(fmt.Sprintf("Service Child %d-%d", i+1, d)).
					WithCategory(gofakeit.RandomString([]string{"Security", "Infrastructure", "Monitoring"})).
					WithCharacteristic("version", gofakeit.AppVersion()).
					WithRelationshipTo(*levelServices[count-1].ID, *levelServices[count-1].Name, "reliesOn").
					Build()
			}
			count++
			levelServices = append(levelServices, child)

		}

		rootWithRel := factory.NewServiceFactory().
			WithName(name).
			WithCategory(gofakeit.RandomString([]string{"Security", "Infrastructure", "Monitoring"})).
			WithCharacteristic("bandwidth", "100").
			WithCharacteristic("unit", "Mbps").
			WithRelationshipTo(*levelServices[count-1].ID, *levelServices[count-1].Name, "reliesOn").
			Build()

		// append levelServices im docs
		for _, s := range levelServices {
			docs = append(docs, s)
		}
		docs = append(docs, rootWithRel)
	}

	_, err := collection.InsertMany(ctx, docs)
	return err
}
