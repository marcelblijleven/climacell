package climacell

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var nowcastEndpoint = "/v3/weather/nowcast"
var unavailableNowcastFields = []field{
	PrecipitationProbability,
	PrecipitationAccumulation,
	CloudSatellite,
	MoonPhase,
	WeatherGroups,
	FireIndex,
}

func (c *Client) Nowcast(
	latitude, longitude float64, unit unit, timeStep int, startTime, endTime *time.Time, fields ...field,
) ([]NowcastData, error) {
	err := validateNowcastArgs(latitude, longitude, startTime, endTime, fields...)

	if err != nil {
		return nil, err
	}

	var startTimeString string
	var endTimeString string

	if startTime == nil {
		startTimeString = "now"
	} else {
		startTime.Format(time.RFC3339)
	}

	if endTime == nil {
		endTimeString = ""
	} else {
		endTime.Format(time.RFC3339)
	}

	var parameters = map[string]interface{}{
		"lat":         latitude,
		"lon":         longitude,
		"unit_system": unit,
		"timestep":    timeStep,
		"start_time":  startTimeString,
		"end_time":    endTimeString,
		"fields":      joinFields(fields, ","),
	}

	req, err := c.makeRequest(nowcastEndpoint, parameters)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkHTTPError(resp, nowcastEndpoint)

	if err != nil {
		return nil, err
	}

	var nowcastData []NowcastData

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&nowcastData); err != nil {
		return nil, err
	}

	return nowcastData, nil
}

func validateNowcastArgs(latitude, longitude float64, startTime, endTime *time.Time, fields ...field) error {
	if !validLatitude(latitude) {
		return ErrInvalidLatitude
	}

	if !validLongitude(longitude) {
		return ErrInvalidLongitude
	}

	var invalidNowcastFields []field

	for _, providedField := range fields {
		for _, unavailableField := range unavailableNowcastFields {
			if providedField == unavailableField {
				invalidNowcastFields = append(invalidNowcastFields, providedField)
			}
		}
	}

	if startTime != nil && endTime != nil {
		start := *startTime
		end := *endTime
		if start.After(end) {
			return errors.New("startTime must be before endTime")
		}
	}

	if invalidNowcastFields != nil {
		msg := joinFields(invalidNowcastFields, ", ")
		return newBadRequestError(nowcastEndpoint, fmt.Sprintf("invalid fields provided (%v)", msg))
	}

	return nil
}
