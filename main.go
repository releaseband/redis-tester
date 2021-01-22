package main

import (
	"log"

	"github.com/releaseband/redis-tester/cli/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
