package climacell

import "testing"

func TestUnit_String(t *testing.T) {
	tests := []struct {
		name string
		u    unit
		want string
	}{
		{
			name: "test String() for unit Si",
			u:    Si,
			want: "si",
		},
		{
			name: "test String() for unit Us",
			u:    Us,
			want: "us",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
