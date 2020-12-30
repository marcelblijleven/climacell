package climacell

import (
	"encoding/json"
	"fmt"
)

var realtimeEndpoint = "/v3/weather/realtime"
var unavailableRealtimeFields = []field{
	PrecipitationProbability,
	PrecipitationAccumulation,
	CloudSatellite,
	WeatherGroups,
}

// Realtime calls the realtime climacell endpoint with the provided fields
func (c *Client) Realtime(latitude, longitude float64, unit unit, fields ...field) (*RealtimeData, error) {
	err := validateRealtimeArgs(latitude, longitude, fields...)

	if err != nil {
		return nil, err
	}

	var parameters = map[string]interface{}{
		"lat":         latitude,
		"lon":         longitude,
		"unit_system": unit,
		"fields":      joinFields(fields, ","),
	}

	req, err := c.makeRequest(realtimeEndpoint, parameters)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkHTTPError(resp, realtimeEndpoint)

	if err != nil {
		return nil, err
	}

	var realtimeData RealtimeData

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&realtimeData); err != nil {
		return nil, err
	}

	return &realtimeData, nil
}

func validateRealtimeArgs(latitude, longitude float64, fields ...field) error {
	if !validLatitude(latitude) {
		return ErrInvalidLatitude
	}

	if !validLongitude(longitude) {
		return ErrInvalidLongitude
	}

	var invalidRealtimeFields []field

	for _, providedField := range fields {
		for _, unavailableField := range unavailableRealtimeFields {
			if providedField == unavailableField {
				invalidRealtimeFields = append(invalidRealtimeFields, providedField)
			}
		}
	}

	if invalidRealtimeFields != nil {
		msg := joinFields(invalidRealtimeFields, ", ")
		return newBadRequestError(realtimeEndpoint, fmt.Sprintf("invalid fields provided (%v)", msg))
	}

	return nil
}
