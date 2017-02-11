package dockerstats

import (
	"errors"
	"testing"
)

func TestCurrent(t *testing.T) {
	tests := []struct {
		expected StatsResult
	}{
		{StatsResult{make([]Stats, 5), nil}},
		{StatsResult{nil, errors.New("test error")}},
	}

	for _, tt := range tests {
		defaultComm := DefaultCommunicator
		DefaultCommunicator = testComm{func() ([]Stats, error) {
			return tt.expected.Stats, tt.expected.Error
		}}
		defer func() { DefaultCommunicator = defaultComm }()

		s, err := Current()
		if err != tt.expected.Error {
			t.Fatalf("Unexpected Error, expected=%v, got=%v", tt.expected.Error, err)
		}

		if len(s) != len(tt.expected.Stats) {
			t.Fatalf("Unexpected Stats, expected=%v, got=%v", tt.expected.Stats, s)
		}
	}
}
