package flags

import (
	"testing"
)

func TestParseLoopsFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    int
		wantErr bool
	}{
		{"default", []string{"cmd"}, 2000, false},
		{"short flag", []string{"cmd", "-pl=42"}, 42, false},
		{"long flag", []string{"cmd", "--piLoops=99"}, 99, false},
		{"reject -piLoops", []string{"cmd", "-piLoops=5"}, 0, true},
		{"reject --pl", []string{"cmd", "--pl=5"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLoopsFlag(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLoopsFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != nil && *got != tt.want {
				t.Errorf("ParseLoopsFlag() = %v, want %v", *got, tt.want)
			}
		})
	}
}
