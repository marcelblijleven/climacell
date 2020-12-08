package climacell

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_validLatitude(t *testing.T) {
	type args struct {
		latitude float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid latitude",
			args: args{latitude: 52},
			want: true,
		},
		{
			name: "valid latitude lower edge",
			args: args{latitude: -59.9},
			want: true,
		},
		{
			name: "valid latitude upper edge",
			args: args{latitude: 59.9},
			want: true,
		},
		{
			name: "invalid latitude lower edge",
			args: args{latitude: -59.91},
			want: false,
		},
		{
			name: "invalid latitude upper edge",
			args: args{latitude: 59.91},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validLatitude(tt.args.latitude); got != tt.want {
				t.Errorf("validLatitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validLongitude(t *testing.T) {
	type args struct {
		longitude float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid longitude",
			args: args{longitude: 100},
			want: true,
		},
		{
			name: "valid longitude lower edge",
			args: args{longitude: -180},
			want: true,
		},
		{
			name: "valid longitude upper edge",
			args: args{longitude: 180},
			want: true,
		},
		{
			name: "invalid longitude lower edge",
			args: args{longitude: -180.01},
			want: false,
		},
		{
			name: "invalid longitude upper edge",
			args: args{longitude: 180.01},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validLongitude(tt.args.longitude); got != tt.want {
				t.Errorf("validLongitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURL(t *testing.T) {
	baseURL := "http://localhost"
	endpoint := "/weather/test"
	u, err := getURL(baseURL, endpoint)

	if err != nil {
		t.Errorf("getUrl() error = %v, expected nil", err)
		return
	}

	want := "http://localhost/weather/test"

	if u.String() != want {
		t.Errorf("getUrl() got = %v, want %v", u.String(), want)
	}
}

func Test_getURL_incorrectBaseURL(t *testing.T) {
	baseURL := " http://localhost"
	endpoint := ""

	_, err := getURL(baseURL, endpoint)
	assert.Error(t, err)
}

func Test_floatToString(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Two digits",
			args: args{number: 52.12},
			want: "52.12",
		},
		{
			name: "Ten digits",
			args: args{number: 52.1234567891},
			want: "52.1234567891",
		},
		{
			name: "Latitude test",
			args: args{number: 52.5579234919951},
			want: "52.5579234919951",
		},
		{
			name: "Longitude test",
			args: args{number: 4.894296098279855},
			want: "4.894296098279855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floatToString(tt.args.number); got != tt.want {
				t.Errorf("floatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinFields(t *testing.T) {
	type args struct {
		fields []field
		sep    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single field",
			args: args{
				fields: []field{Temperature},
				sep:    ", ",
			},
			want: "temp",
		},
		{
			name: "multiple fields",
			args: args{
				fields: []field{Temperature, CloudCeiling},
				sep:    ", ",
			},
			want: "temp, cloud_ceiling",
		},
		{
			name: "empty fields",
			args: args{
				fields: []field{},
				sep:    ", ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinFields(tt.args.fields, tt.args.sep); got != tt.want {
				t.Errorf("joinFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHTTPError(t *testing.T) {
	type args struct {
		resp     *http.Response
		endpoint string
	}

	mockBody := struct {
		Message string `json:"message"`
	}{
		Message: "mock body content",
	}

	data, err := json.Marshal(&mockBody)

	if err != nil {
		t.Fatal("error setting up mockBody")
	}

	body := ioutil.NopCloser(bytes.NewBuffer(data))

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "response status code 200",
			args: args{
				resp:     &http.Response{StatusCode: 200},
				endpoint: "/v3/weather_test",
			},
			wantErr: false,
		},
		{
			name: "response status code 400",
			args: args{
				resp: &http.Response{
					StatusCode: 400,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code 401",
			args: args{
				resp: &http.Response{
					StatusCode: 401,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code 403",
			args: args{
				resp: &http.Response{
					StatusCode: 403,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code 404",
			args: args{
				resp: &http.Response{
					StatusCode: 404,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code 429",
			args: args{
				resp: &http.Response{
					StatusCode: 429,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code 500",
			args: args{
				resp: &http.Response{
					StatusCode: 500,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
		{
			name: "response status code not specified in function logic",
			args: args{
				resp: &http.Response{
					StatusCode: 503,
					Body:       body,
				},
				endpoint: "/v3/weather_test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkHTTPError(tt.args.resp, tt.args.endpoint); (err != nil) != tt.wantErr {
				t.Errorf("checkHTTPError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
