package dockerstats

import (
	"sync"
)

type testComm struct {
	statsFn func() ([]Stats, error)
}

func (t testComm) Stats() ([]Stats, error) {
	return t.statsFn()
}

// Monitor provides the ability to recieve a constant stream of statistics for each running Docker
// container.
//
// Each `StatsResult` sent through the channel contains either an `error` or a
// `Stats` slice equal in length to the number of running Docker containers.
type Monitor struct {
	Stream chan *StatsResult
	Comm   Communicator

	mu      sync.Mutex
	stopped bool
}

// NewMonitor initializes and returns a Monitor which can recieve a stream of Docker container statistics.
func NewMonitor() *Monitor {
	m := Monitor{
		Stream: make(chan *StatsResult),
		Comm:   DefaultCommunicator,
	}
	m.start()

	return &m
}

// Stop tells the monitor to stop streaming Docker container statistics.
func (m *Monitor) Stop() {
	m.mu.Lock()
	m.stopped = true
	m.mu.Unlock()
}

// start begins polling for Docker container statistics, and sends them through the Monitor's
// stream to be consumed.
func (m *Monitor) start() {
	go func() {
		for {
			m.mu.Lock()
			stopped := m.stopped
			// Do not defer! If the channel blocks below it
			// can lead to deadlock situations.
			m.mu.Unlock()

			if stopped {
				close(m.Stream)
				break
			}

			s, err := m.Comm.Stats()
			m.Stream <- &StatsResult{
				Stats: s,
				Error: err,
			}
		}
	}()
}
