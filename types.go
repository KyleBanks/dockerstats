package dockerstats

// StatsResult is the value recieved when using Monitor to listen for
// Docker statistics.
type StatsResult struct {
	Stats []Stats
	Error error
}

// Stats contains the statistics of a currently running Docker container.
type Stats struct {
	Container string
	Memory    MemoryStats
	CPU       string
}

// String returns a human-readable string containing the details of a Stats value.
func (s Stats) String() string {
	return fmt.Sprintf("Container=%v Memory={%v} CPU=%v", s.Container, s.Memory, s.CPU)
}

// MemoryStats contains the statistics of a running Docker container related to
// memory usage.
type MemoryStats struct {
	Raw     string
	Percent string
}

// String returns a human-readable string containing the details of a MemoryStats value.
func (m MemoryStats) String() string {
	return fmt.Sprintf("Raw=%v, Percent=%v", m.Raw, m.Percent)
}
