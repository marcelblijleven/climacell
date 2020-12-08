package climacell

import (
	"encoding/json"
	"testing"
)

func TestFloatData_String(t *testing.T) {
	type fields struct {
		Value *float64
		Units string
	}
	floatValue := 13.3712345
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "non nil value",
			fields: fields{
				Value: &floatValue,
				Units: "cm",
			},
			want: "13.3712345 cm",
		},
		{
			name: "nil value",
			fields: fields{
				Value: nil,
				Units: "cm",
			},
			want: "0 cm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &FloatData{
				Value: tt.fields.Value,
				Units: tt.fields.Units,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntData_String(t *testing.T) {
	type fields struct {
		Value *int
		Units string
	}
	intValue := 13
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "non nil value",
			fields: fields{
				Value: &intValue,
				Units: "cm",
			},
			want: "13 cm",
		},
		{
			name: "nil value",
			fields: fields{
				Value: nil,
				Units: "cm",
			},
			want: "0 cm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &IntData{
				Value: tt.fields.Value,
				Units: tt.fields.Units,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringData_String(t *testing.T) {
	type fields struct {
		Value *string
	}
	stringValue := "weather_test"
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "non nil value",
			fields: fields{
				Value: &stringValue,
			},
			want: "weather_test",
		},
		{
			name: "nil value",
			fields: fields{
				Value: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &StringData{
				Value: tt.fields.Value,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeData_UnmarshalJSON(t *testing.T) {
	data := []byte("{\"value\": \"2020-12-07T20:06:54.764Z\"}")
	var timeData TimeData
	err := json.Unmarshal(data, &timeData)

	if err != nil {
		t.Errorf("TimeData UnmarshalJSON() error = %v, expected nil", err)
	}
}
