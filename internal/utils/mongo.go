package utils

import (
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"
)

func GerFieldsProjection(fieldsParam *string) bson.M {
	if fieldsParam == nil || *fieldsParam == "" {
		return nil // Return an empty map instead of nil
	}

	fields := strings.Split(*fieldsParam, ",")
	projection := bson.M{"id": 1} // Always include ID

	for _, field := range fields {
		projection[field] = 1
	}

	return projection
}

func ConvertBsonMToMinimalJSONResponse(record mongo.SingleResult) (map[string]interface{}, error) {
	var result bson.M
	err := record.Decode(&result)
	if err != nil {
		return nil, err
	}
	// Convert `_id` to string if present
	if id, ok := result["_id"].(primitive.ObjectID); ok {
		result["id"] = id.Hex()
		delete(result, "_id") // Remove original _id key
	}

	// Convert `primitive.DateTime` fields to `strfmt.DateTime`
	for key, value := range result {
		if dt, ok := value.(primitive.DateTime); ok {
			result[key] = strfmt.DateTime(dt.Time()) // Convert to expected format
		}
	}

	// Convert `bson.M` to a minimal JSON response
	response := make(map[string]interface{})
	for key, value := range result {
		response[key] = value
	}

	return response, nil
}

func BuildTmfMongoFilter(queryParams map[string][]string, usePipeline bool) (interface{}, bool) {
	// Exclude pagination & projection params
	delete(queryParams, "fields")
	delete(queryParams, "limit")
	delete(queryParams, "offset")

	// Check for deep parameter
	//depth := -1
	//if deepVals, ok := queryParams["deep"]; ok && len(deepVals) > 0 {
	//	if d, err := strconv.Atoi(deepVals[0]); err == nil {
	//		depth = d
	//	}
	//	delete(queryParams, "deep")
	//}

	filter := bson.M{}
	orFilters := []bson.M{}

	opMap := map[string]string{
		"eq":  "$eq",
		"gt":  "$gt",
		"gte": "$gte",
		"lt":  "$lt",
		"lte": "$lte",
		"ne":  "$ne",
		"in":  "$in",
		"nin": "$nin",
	}

	for key, values := range queryParams {
		if len(values) == 0 {
			continue
		}

		parts := strings.Split(key, ".")
		filterOperator := parts[len(parts)-1]
		filterField := key
		operator := "$eq"

		if len(parts) > 1 {
			if op, exists := opMap[filterOperator]; exists {
				operator = op
				filterField = strings.Join(parts[:len(parts)-1], ".")
			}
		}

		rawValue := values[0]
		multipleValues := strings.Split(rawValue, ",")

		if operator == "$in" || operator == "$nin" {
			var parsedItems []interface{}
			for _, v := range multipleValues {
				parsedItems = append(parsedItems, parseValue(v))
			}
			filter[filterField] = bson.M{operator: parsedItems}
			continue
		}

		if len(multipleValues) > 1 {
			for _, v := range multipleValues {
				orFilters = append(orFilters, bson.M{
					filterField: bson.M{operator: parseValue(v)},
				})
			}
			continue
		}

		filter[filterField] = bson.M{operator: parseValue(rawValue)}
	}

	if len(orFilters) > 0 {
		if existingOr, exists := filter["$or"]; exists {
			filter["$or"] = append(existingOr.([]bson.M), orFilters...)
		} else {
			filter["$or"] = orFilters
		}
	}

	if usePipeline == true {
		pipeline := mongo.Pipeline{
			{{Key: "$match", Value: filter}},
		}
		return pipeline, true
	}
	// If deep mode is enabled, return aggregate pipeline
	//if depth >= 0 {
	//	pipeline := mongo.Pipeline{
	//		{{Key: "$match", Value: filter}},
	//		{{Key: "$graphLookup", Value: bson.M{
	//			"from":             "serviceInventory",
	//			"startWith":        "$serviceRelationship.service.id",
	//			"connectFromField": "serviceRelationship.service.id",
	//			"connectToField":   "id",
	//			"as":               "relatedServices",
	//			"depthField":       "level",
	//			"maxDepth":         depth,
	//		}}},
	//	}
	//	return pipeline, true
	//}

	return filter, false
}

func parseValue(value string) interface{} {
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return floatValue
	} else if dateValue, err := time.Parse(time.RFC3339, value); err == nil {
		return dateValue
	}
	return value
}
