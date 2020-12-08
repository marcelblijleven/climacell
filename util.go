package climacell

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func validLatitude(latitude float64) bool {
	return -59.9 <= latitude && latitude <= 59.9
}

func validLongitude(longitude float64) bool {
	return -180 <= longitude && longitude <= 180
}

func getURL(baseURL, endpoint string) (*url.URL, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	e, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	return u.ResolveReference(e), nil
}

func floatToString(number float64) string {
	return fmt.Sprintf("%g", number)
}

func joinFields(fields []field, sep string) string {
	var fieldNames []string

	for _, f := range fields {
		fieldNames = append(fieldNames, f.String())
	}

	return strings.Join(fieldNames, sep)
}

func checkHTTPError(resp *http.Response, endpoint string) error {
	if resp.StatusCode == 200 {
		return nil
	}

	var httpError HTTPError

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&httpError); err != nil {
		return err
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		if resp.StatusCode == 400 {
			// Bad request
			return newBadRequestError(endpoint, httpError.Message)
		}

		if resp.StatusCode == 401 {
			// Unauthorized
			return newUnauthorizedError(endpoint, httpError.Message)
		}

		if resp.StatusCode == 403 {
			// Forbidden
			return newForbiddenError(endpoint, httpError.Message)
		}

		if resp.StatusCode == 404 {
			// Not found
			return newNotFoundError(endpoint, httpError.Message)
		}

		if resp.StatusCode == 429 {
			// Too many requests
			return newTooManyRequestError(endpoint, httpError.Message)
		}

		if resp.StatusCode == 500 {
			// Internal server error
			return newInternalServerError(endpoint, httpError.Message)
		}
	}

	return &httpError
}
