package dockerstats

import (
	"strconv"
	"sync"
	"testing"
)

func TestNewMonitor(t *testing.T) {
	maxRecieves := 5
	var receiveCount int

	m := NewMonitor()
	m.Comm = testComm{func() ([]Stats, error) {
		s := []Stats{
			{Container: strconv.Itoa(receiveCount)},
		}
		receiveCount++
		return s, nil
	}}

	for i := 0; i < maxRecieves; i++ {
		s := <-m.Stream

		if s.Error != nil {
			t.Fatal(s.Error)
		}

		if len(s.Stats) != 1 {
			t.Fatalf("Unexpected number of stats recieved on iteration %v, expected=1, got=%v.\n%v", i, len(s.Stats), s)
		} else if s.Stats[0].Container != strconv.Itoa(i) {
			t.Fatalf("Unexpected container stats recieved on iteration %v, expected=%v, got=%v.\n%v", i, i, s.Stats[0].Container, s)
		}
	}
}

func TestMonitor_Stop(t *testing.T) {
	var mu sync.Mutex

	m := NewMonitor()
	var callCount int
	m.Comm = testComm{func() ([]Stats, error) {
		mu.Lock()
		callCount++
		mu.Unlock()

		return make([]Stats, 0), nil
	}}

	s := <-m.Stream
	if s.Error != nil {
		t.Fatal(s.Error)
	}

	mu.Lock()
	expected := callCount
	mu.Unlock()
	m.Stop()

	// Read the values from the channel and ensure at the end that the channel
	// has been closed.
	for i := 0; i <= expected-1; i++ {
		select {
		case _, ok := <-m.Stream:
			if ok && i == expected-1 {
				t.Fatal("Expected stream to be closed after the last record is read")
			}
		}
	}

	// Ensure the callCount has stopped increasing.
	if expected != callCount {
		t.Fatalf("Unexpected callCount, Monitor should have stopped, expected=%v, got=%v", expected, callCount)
	}
}
