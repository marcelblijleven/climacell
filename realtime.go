package climacell

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	u, err := getURL(c.baseURL, realtimeEndpoint)

	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("lat", floatToString(latitude))
	q.Set("lon", floatToString(longitude))
	q.Set("unit_system", unit.String())
	q.Set("fields", joinFields(fields, ","))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", c.apiKey)

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

	if err := json.NewDecoder(resp.Body).Decode(&realtimeData); err != nil {
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
