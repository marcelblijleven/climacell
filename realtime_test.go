package climacell

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_validateRealtimeArgs(t *testing.T) {
	type args struct {
		latitude  float64
		longitude float64
		fields    []field
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid arguments",
			args: args{
				latitude:  59.9,
				longitude: 180,
				fields: []field{
					Temperature,
					FeelsLike,
					Precipitation,
				},
			},
		},
		{
			name: "invalid latitude",
			args: args{
				latitude:  59.91,
				longitude: 180,
				fields: []field{
					Temperature,
					FeelsLike,
					Precipitation,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid longitude",
			args: args{
				latitude:  59.9,
				longitude: 181,
				fields: []field{
					Temperature,
					FeelsLike,
					Precipitation,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid field",
			args: args{
				latitude:  59.9,
				longitude: 180,
				fields: []field{
					Temperature,
					FeelsLike,
					Precipitation,
					WeatherGroups,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRealtimeArgs(tt.args.latitude, tt.args.longitude, tt.args.fields...); (err != nil) != tt.wantErr {
				t.Errorf("validateRealtimeArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateRealtimeArgs_singleFieldError(t *testing.T) {
	err := validateRealtimeArgs(59.9, 180, WeatherGroups, Temperature)

	if err == nil {
		t.Error("validateRealtimeArgs() error, expected err to be non nil")
		return
	}

	want := "bad request /v3/weather/realtime: invalid fields provided (weather_groups)"
	if err.Error() != want {
		t.Errorf("validateRealtimeArgs() error, got = %q, want %q", err.Error(), want)
	}
}

func Test_validateRealtimeArgs_multipleFieldError(t *testing.T) {
	err := validateRealtimeArgs(59.9, 180,
		PrecipitationProbability,
		PrecipitationAccumulation,
		CloudSatellite,
		WeatherGroups,
	)

	if err == nil {
		t.Error("validateRealtimeArgs() error, expected err to be non nil")
		return
	}

	want := "bad request /v3/weather/realtime: invalid fields provided (precipitation_probability, precipitation_accumulation, cloud_satellite, weather_groups)"
	if err.Error() != want {
		t.Errorf("validateRealtimeArgs() error, got = %q, want %q", err.Error(), want)
	}
}

func TestClient_Realtime(t *testing.T) {
	mockData, err := ioutil.ReadFile("mocks/realtime_response.json")

	if err != nil {
		t.Fatal("error setting up mock response file")
		return
	}

	srv, closeFunc := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(mockData)
	})

	defer closeFunc()

	c, err := NewClient("apikey", srv.Client())

	if err != nil {
		t.Fatal("error setting up client")
		return
	}

	latitude := 52.321234567890
	longitude := 4.95124567890

	resp, err := c.Realtime(latitude, longitude, Si, Temperature)

	if err != nil {
		t.Errorf("Realtime() error = %v, want nil", err.Error())
		return
	}

	assert.Equal(t, 3.63, *resp.Temperature.Value)
}

func TestClient_Realtime_non200response(t *testing.T) {
	mockError := HTTPError{
		Endpoint: "/v3/weather/realtime",
		Message:  "mock error message",
	}
	mockData, err := json.Marshal(mockError)

	if err != nil {
		t.Fatal("error setting up mock body")
	}

	srv, closeFunc := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write(mockData)
	})

	defer closeFunc()

	c, err := NewClient("apikey", srv.Client())

	if err != nil {
		t.Fatal("error setting up client")
	}

	latitude := 52.321234567890
	longitude := 4.95124567890

	resp, err := c.Realtime(latitude, longitude, Si, Temperature)

	if err == nil {
		t.Error("Realtime() error = nil, expected non nil")
		return
	}

	errMsg := err.Error()
	wantMsg := "bad request /v3/weather/realtime: mock error message"

	if errMsg != wantMsg {
		t.Errorf("Realtime() error message = %q, want %q", errMsg, wantMsg)
	}

	assert.Nil(t, resp)
}

func TestClient_Realtime_invalidArgs(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
		apiKey     string
	}
	type args struct {
		latitude  float64
		longitude float64
		unit      unit
		fields    []field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RealtimeData
		wantErr bool
	}{
		{
			name: "invalid latitude",
			fields: fields{
				httpClient: nil,
				baseURL:    "",
				apiKey:     "apikey",
			},
			args: args{
				latitude:  -59.91,
				longitude: 180,
				unit:      Si,
				fields:    nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid longitude",
			fields: fields{
				httpClient: nil,
				baseURL:    "",
				apiKey:     "apikey",
			},
			args: args{
				latitude:  -59.9,
				longitude: 180.01,
				unit:      Si,
				fields:    nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid base URL",
			fields: fields{
				httpClient: nil,
				baseURL:    " http://localhost",
				apiKey:     "apikey",
			},
			args: args{
				latitude:  -59.9,
				longitude: 180,
				unit:      Si,
				fields:    nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				baseURL:    tt.fields.baseURL,
				apiKey:     tt.fields.apiKey,
			}
			got, err := c.Realtime(tt.args.latitude, tt.args.longitude, tt.args.unit, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Realtime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Realtime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
