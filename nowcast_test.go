package climacell

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_validateNowcastArgs(t *testing.T) {
	type args struct {
		latitude  float64
		longitude float64
		startTime *time.Time
		endTime   *time.Time
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
			if err := validateNowcastArgs(tt.args.latitude, tt.args.longitude, tt.args.startTime, tt.args.endTime, tt.args.fields...); (err != nil) != tt.wantErr {
				t.Errorf("validateNowcastArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateNowcastArgs_singleFieldError(t *testing.T) {
	err := validateNowcastArgs(59.9, 180, nil, nil, WeatherGroups, Temperature)

	if err == nil {
		t.Error("validateNowcastArgs() error, expected err to be non nil")
		return
	}

	want := "bad request /v3/weather/nowcast: invalid fields provided (weather_groups)"
	if err.Error() != want {
		t.Errorf("validateNowcastArgs() error, got = %q, want %q", err.Error(), want)
	}
}

func Test_validateNowcastArgs_multipleFieldError(t *testing.T) {
	err := validateNowcastArgs(59.9, 180, nil, nil,
		PrecipitationProbability,
		PrecipitationAccumulation,
		CloudSatellite,
		WeatherGroups,
	)

	if err == nil {
		t.Error("validateNowcastArgs() error, expected err to be non nil")
		return
	}

	want := "bad request /v3/weather/nowcast: invalid fields provided (precipitation_probability, precipitation_accumulation, cloud_satellite, weather_groups)"
	if err.Error() != want {
		t.Errorf("validateNowcastArgs() error, got = %q, want %q", err.Error(), want)
	}
}

func TestClient_Nowcast(t *testing.T) {
	mockData, err := ioutil.ReadFile("mocks/nowcast_response.json")

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
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	startTime := time.Date(2020, 12, 31, 4, 55, 35, 176, loc)

	resp, err := c.Nowcast(latitude, longitude, Si, 5, &startTime, nil,
		Temperature,
		FeelsLike,
		DewPoint,
		BarometricPressure,
		CloudBase,
		CloudCeiling,
	)

	if err != nil {
		t.Errorf("Nowcast() error = %v, want nil", err.Error())
		return
	}

	assert.Equal(t, 3.63, *resp[0].Temperature.Value)
}
