package main

import (
	"github.com/ayupov-ayaz/redis-tester/cli/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
