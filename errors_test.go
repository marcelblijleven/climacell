package climacell

import (
	"fmt"
	"testing"
)

func TestHTTPError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := &HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}
	want := fmt.Sprintf("error %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestBadRequestError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newBadRequestError(endpoint, message)
	want := fmt.Sprintf("bad request %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestForbiddenError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newForbiddenError(endpoint, message)
	want := fmt.Sprintf("forbidden %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestInternalServerError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newInternalServerError(endpoint, message)
	want := fmt.Sprintf("internal server error %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestTooManyRequestError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newTooManyRequestError(endpoint, message)
	want := fmt.Sprintf("too many requests %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestUnauthorizedError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newUnauthorizedError(endpoint, message)
	want := fmt.Sprintf("unauthorized %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestNotFoundError_Error(t *testing.T) {
	endpoint := "/v3/weather/test"
	message := "test message"
	e := newNotFoundError(endpoint, message)
	want := fmt.Sprintf("not found error %v: %v", endpoint, message)
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}
