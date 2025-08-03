package factories

import (
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"time"
)

type ServiceFactory struct {
	service *models.Service
}

func init() {
	gofakeit.Seed(0)
}

// NewServiceFactory Constructor
func NewServiceFactory() *ServiceFactory {
	id := uuid.New().String()
	return &ServiceFactory{
		service: &models.Service{
			ID:          &id,
			Name:        ptr(gofakeit.AppName()),
			Description: ptr(gofakeit.Sentence(10)),
			StartDate:   ptr(toStrfmtDateTime(gofakeit.Date())),
			EndDate:     ptr(toStrfmtDateTime(gofakeit.Date())),
		},
	}
}

func CreateMany(count int, builder func(i int) *ServiceFactory) []*models.Service {
	services := make([]*models.Service, 0, count)
	for i := 0; i < count; i++ {
		services = append(services, builder(i).Build())
	}
	return services
}

// Chainable
func (f *ServiceFactory) WithName(name string) *ServiceFactory {
	f.service.Name = &name
	return f
}

func (f *ServiceFactory) WithCategory(category string) *ServiceFactory {
	f.service.Category = &category
	return f
}

func (f *ServiceFactory) WithDescription(desc string) *ServiceFactory {
	f.service.Description = &desc
	return f
}

func (f *ServiceFactory) WithID(id string) *ServiceFactory {
	f.service.ID = &id
	return f
}

func (f *ServiceFactory) WithStartDate(t time.Time) *ServiceFactory {
	f.service.StartDate = ptr(toStrfmtDateTime(t))
	return f
}

func (f *ServiceFactory) WithRelationshipTo(serviceID string, name string, relationType string) *ServiceFactory {
	f.service.ServiceRelationship = append(f.service.ServiceRelationship, &models.ServiceRelationship{
		RelationshipType: relationType,
		Service: &models.ServiceRefOrValue{
			ID:   &serviceID,
			Name: &name,
		},
	})
	return f
}

func (f *ServiceFactory) WithRelatedParty(role string, name string) *ServiceFactory {
	f.service.RelatedParty = append(f.service.RelatedParty, &models.RelatedParty{
		Role: ptr(role),
		Name: ptr(name),
	})
	return f
}

func (f *ServiceFactory) WithCharacteristic(name, value string) *ServiceFactory {
	f.service.ServiceCharacteristic = append(f.service.ServiceCharacteristic, &models.Characteristic{
		Name:  name,
		Value: value,
	})
	return f
}

func (f *ServiceFactory) Build() *models.Service {
	return f.service
}

// --- Helpers ---
func ptr[T any](v T) *T {
	return &v
}

func toStrfmtDateTime(t time.Time) strfmt.DateTime {
	return strfmt.DateTime(t)
}
