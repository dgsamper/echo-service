package echo_test

import (
	"testing"

	"github.com/dgsamper/echo-service/internal/echo"
)

func TestParsePort(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    string
		wantErr bool
	}{
		{name: "default", value: "", want: "8080"},
		{name: "custom", value: "9090", want: "9090"},
		{name: "not a number", value: "http", wantErr: true},
		{name: "below range", value: "0", wantErr: true},
		{name: "above range", value: "65536", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := echo.ParsePort(tt.value)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParsePort(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParsePort(%q) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}
