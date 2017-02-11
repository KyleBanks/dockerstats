package dockerstats

import (
	"encoding/json"
	"os/exec"
	"strings"
)

// Communicator provides an interface for communicating with and retrieving stats
// from Docker.
type Communicator interface {
	Stats() ([]Stats, error)
}

// CliCommunicator uses the Docker CLI to retrieve stats for currently running Docker
// containers.
type CliCommunicator struct {
	DockerPath string
	Command    []string
}

// Stats returns Docker container statistics using the Docker CLI.
func (c CliCommunicator) Stats() ([]Stats, error) {
	out, err := exec.Command(c.DockerPath, c.Command...).Output()
	if err != nil {
		return nil, err
	}

	containers := strings.Split(string(out), "\n")
	stats := make([]Stats, 0)
	for _, con := range containers {
		if len(con) == 0 {
			continue
		}

		var s Stats
		if err := json.Unmarshal([]byte(con), &s); err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
