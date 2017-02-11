// Package main is a sample use of dockerstats that infinitely listens for Docker stats,
// and writes the output to a file.
//
// Try it out like so:
//
// 		go run main.go output.txt
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KyleBanks/dockerstats"
)

func main() {
	filename := "output.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.Printf("Writing output to '%v'", f.Name())

	m := dockerstats.NewMonitor()
	for res := range m.Stream {
		if res.Error != nil {
			panic(err)
		}

		if len(res.Stats) == 0 {
			log.Println("No Docker containers running, output complete.")
			m.Stop()
			break
		}

		var out string
		for _, s := range res.Stats {
			out = fmt.Sprintf("%v%v: %v\n", out, time.Now(), s)
		}

		if _, err := f.WriteString(out); err != nil {
			panic(err)
		}
	}
}
