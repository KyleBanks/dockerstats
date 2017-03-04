# dockerstats
--
    import "github.com/KyleBanks/dockerstats"

[![Build Status](https://travis-ci.org/KyleBanks/dockerstats.svg?branch=master)](https://travis-ci.org/KyleBanks/dockerstats) &nbsp;
[![GoDoc](https://godoc.org/github.com/KyleBanks/dockerstats?status.svg)](https://godoc.org/github.com/KyleBanks/dockerstats) &nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/dockerstats)](https://goreportcard.com/report/github.com/KyleBanks/dockerstats)

Package dockerstats provides the ability to get currently running Docker
container statistics, including memory and CPU usage.

To get the statistics of running Docker containers, you can use the `Current()`
function:

    stats, err := dockerstats.Current()
    if err != nil {
    	panic(err)
    }

    for _, s := range stats {
    	fmt.Println(s.Container) // 9f2656020722
    	fmt.Println(s.Memory) // {Raw=221.7 MiB / 7.787 GiB, Percent=2.78%}
    	fmt.Println(s.CPU) // 99.79%
    }

Alternatively, you can use the `NewMonitor()` function to receive a constant
stream of Docker container stats, available on the Monitor's `Stream` channel:

    m := dockerstats.NewMonitor()

    for res := range m.Stream {
    	if res.Error != nil {
    		panic(err)
    	}

    	for _, s := range res.Stats {
    		fmt.Println(s.Container) // 9f2656020722
    	}
    }

## Example

For a simple example of writing the statistics of each running Docker container to a file, see the [example directory](./example).

```
cd example
go run main.go "output.txt"
```

## Usage

#### func  Current

```go
func Current() ([]Stats, error)
```
Current returns the current `Stats` of each running Docker container.

Current will always return a `[]Stats` slice equal in length to the number of
running Docker containers, or an `error`. No error is returned if there are no
running Docker containers, simply an empty slice.

#### type CliCommunicator

```go
type CliCommunicator struct {
	DockerPath string
	Command    []string
}
```

CliCommunicator uses the Docker CLI to retrieve stats for currently running
Docker containers.

#### func (CliCommunicator) Stats

```go
func (c CliCommunicator) Stats() ([]Stats, error)
```
Stats returns Docker container statistics using the Docker CLI.

#### type Communicator

```go
type Communicator interface {
	Stats() ([]Stats, error)
}
```

Communicator provides an interface for communicating with and retrieving stats
from Docker.

```go
var DefaultCommunicator Communicator = CliCommunicator{
	DockerPath: defaultDockerPath,
	Command:    []string{defaultDockerCommand, defaultDockerNoStreamArg, defaultDockerFormatArg, defaultDockerFormat},
}
```
DefaultCommunicator is the default way of retrieving stats from Docker.

When calling `Current()`, the `DefaultCommunicator` is used, and when retriving
a `Monitor` using `NewMonitor()`, it is initialized with the
`DefaultCommunicator`.

#### type IOStats

```go
type IOStats struct {
	Network string `json:"network"`
	Block   string `json:"block"`
}
```

IOStats contains the statistics of a running Docker container related to IO,
including network and block.

#### func (IOStats) String

```go
func (i IOStats) String() string
```
String returns a human-readable string containing the details of a IOStats
value.

#### type MemoryStats

```go
type MemoryStats struct {
	Raw     string `json:"raw"`
	Percent string `json:"percent"`
}
```

MemoryStats contains the statistics of a running Docker container related to
memory usage.

#### func (MemoryStats) String

```go
func (m MemoryStats) String() string
```
String returns a human-readable string containing the details of a MemoryStats
value.

#### type Monitor

```go
type Monitor struct {
	Stream chan *StatsResult
	Comm   Communicator
}
```

Monitor provides the ability to recieve a constant stream of statistics for each
running Docker container.

Each `StatsResult` sent through the channel contains either an `error` or a
`Stats` slice equal in length to the number of running Docker containers.

#### func  NewMonitor

```go
func NewMonitor() *Monitor
```
NewMonitor initializes and returns a Monitor which can recieve a stream of
Docker container statistics.

#### func (*Monitor) Stop

```go
func (m *Monitor) Stop()
```
Stop tells the monitor to stop streaming Docker container statistics.

#### type Stats

```go
type Stats struct {
	Container string      `json:"container"`
	Memory    MemoryStats `json:"memory"`
	CPU       string      `json:"cpu"`
	IO        IOStats     `json:"io"`
	PIDs      int         `json:"pids"`
}
```

Stats contains the statistics of a currently running Docker container.

#### func (Stats) String

```go
func (s Stats) String() string
```
String returns a human-readable string containing the details of a Stats value.

#### type StatsResult

```go
type StatsResult struct {
	Stats []Stats `json:"stats"`
	Error error   `json:"error"`
}
```

StatsResult is the value recieved when using Monitor to listen for Docker
statistics.

## License

`dockerstats` is available [under MIT license](./LICENSE).
