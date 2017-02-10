package main

import (
	"log"

	"github.com/KyleBanks/dockerstats"
)

func main() {
	c := dockerstats.Monitor()

	for {
		res := <-c
		log.Println(res)
	}
}
