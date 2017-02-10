// Package dockerstats provides the ability to get currently running Docker container statistics,
// including memory and CPU usage.
//
// To get the statistics of running Docker containers, you can use the `Current()` function:
//
// 		stats, err := dockerstats.Current()
//		if err != nil {
//			panic(err)
//		}
//
//		for _, s := range stats {
//			fmt.Println(s.Container) // 9f2656020722
//			fmt.Println(s.Memory) // {Raw=221.7 MiB / 7.787 GiB, Percent=2.78%}
//			fmt.Println(s.CPU) // 99.79%
//		}
//
// Alternatively, you can use the `Monitor()` function to receive a constant stream of Docker container stats:
//
// 		c := make(chan *StatsResult)
// 		dockerstats.Monitor(c)
//
// 		for {
// 			res := <-c
//			if res.Error != nil {
//				panic(err)
//			}
//
//			for _, con := range res.Stats {
//				fmt.Println(con.Container) // 9f2656020722
//			}
// 		}
package dockerstats

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

const (
	dockerPath        string = "/usr/local/bin/docker"
	dockerCommand     string = "stats"
	dockerNoStreamArg string = "--no-stream"
	dockerFormatArg   string = "--format"
	dockerFormat      string = `{"container":"{{ .Container }}","memory":{"raw":"{{ .MemUsage }}","percent":"{{ .MemPerc }}"},"cpu":"{{ .CPUPerc }}"}`
)

// Monitor repeatedly retrieves the current stats for each running Docker container,
// and sends them through the channel provided.
func Monitor(c chan *StatsResult) {
	go func() {
		for {
			s, err := CurStats()
			c <- &StatsResult{
				Stats: s,
				Error: err,
			}
		}
	}()
}

// Current returns the current stats of each running Docker container.
func Current() ([]Stats, error) {
	out, err := exec.Command(dockerPath, dockerCommand, dockerNoStreamArg, dockerFormatArg, dockerFormat).Output()
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
