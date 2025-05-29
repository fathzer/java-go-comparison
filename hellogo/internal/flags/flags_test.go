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
			got, err := ParseIntFlag("pl", "piLoops", 2000, tt.args)
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

func TestParseStringFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{"default", []string{"cmd"}, "default-value", false},
		{"short flag", []string{"cmd", "-f=test"}, "test", false},
		{"long flag", []string{"cmd", "--fen=rnbqkbnr/8/8/8/8/8/8/RNBQKBNR"}, "rnbqkbnr/8/8/8/8/8/8/RNBQKBNR", false},
		{"reject -fen", []string{"cmd", "-fen=test"}, "", true},
		{"reject --f", []string{"cmd", "--f=test"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringFlag("f", "fen", "default-value", tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != nil && *got != tt.want {
				t.Errorf("ParseStringFlag() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestParseBoolFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    bool
		wantErr bool
	}{
		{"default", []string{"cmd"}, false, false},
		{"short flag", []string{"cmd", "-b"}, true, false},
		{"long flag", []string{"cmd", "--blackPlaying"}, true, false},
		{"reject -blackPlaying", []string{"cmd", "-blackPlaying"}, false, true},
		{"reject --b", []string{"cmd", "--b"}, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBoolFlag("b", "blackPlaying", tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBoolFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseBoolFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
