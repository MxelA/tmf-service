package validator

import (
	"fmt"
	repository "github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/repositories"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/restapi/operations/service"
)

func ValidateBusinessRules(createServiceParams service.CreateServiceParams, repo repository.ServiceInventoryRepository) []*models.Error {
	var errs []*models.Error
	
	for _, serviceRelation := range createServiceParams.Service.ServiceRelationship {
		if serviceRelation.Service != nil && serviceRelation.Service.ID != nil && *serviceRelation.Service.ID != "" {
			serviceInventory, err := repo.GetByID(createServiceParams.HTTPRequest.Context(), *serviceRelation.Service.ID, nil)

			if err != nil {
				code := "500"
				reason := fmt.Sprintf("Error checking ExternalIdentifier serviceRelation.Service.Id: %v", err)
				errs = append(errs, &models.Error{
					Code:    &code,
					Reason:  &reason,
					Message: "Internal validation error",
				})
			} else {
				serviceRelation.Service.ID = serviceInventory.ID
			}
		}
	}

	return errs
}
