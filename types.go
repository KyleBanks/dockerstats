package dockerstats

import (
	"fmt"
)

// Stats contains the statistics of a currently running Docker container.
type Stats struct {
	Container string      `json:"container"`
	Memory    MemoryStats `json:"memory"`
	CPU       string      `json:"cpu"`
	IO        IOStats     `json:"io"`
	PIDs      int         `json:"pids"`
}

// String returns a human-readable string containing the details of a Stats value.
func (s Stats) String() string {
	return fmt.Sprintf("Container=%v Memory={%v} CPU=%v IO={%v} PIDs=%v", s.Container, s.Memory, s.CPU, s.IO, s.PIDs)
}

// MemoryStats contains the statistics of a running Docker container related to
// memory usage.
type MemoryStats struct {
	Raw     string `json:"raw"`
	Percent string `json:"percent"`
}

// String returns a human-readable string containing the details of a MemoryStats value.
func (m MemoryStats) String() string {
	return fmt.Sprintf("Raw=%v Percent=%v", m.Raw, m.Percent)
}

// IOStats contains the statistics of a running Docker container related to
// IO, including network and block.
type IOStats struct {
	Network string `json:"network"`
	Block   string `json:"block"`
}

// String returns a human-readable string containing the details of a IOStats value.
func (i IOStats) String() string {
	return fmt.Sprintf("Network=%v Block=%v", i.Network, i.Block)
}

// StatsResult is the value recieved when using Monitor to listen for
// Docker statistics.
type StatsResult struct {
	Stats []Stats `json:"stats"`
	Error error   `json:"error"`
}
