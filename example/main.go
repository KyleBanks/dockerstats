package main

import (
	"log"

	"github.com/KyleBanks/dockerstats"
)

func main() {
	c := make(chan *dockerstats.StatsResult)
	dockerstats.Monitor(c)

	for {
		res := <-c
		log.Println(res)
	}
}
