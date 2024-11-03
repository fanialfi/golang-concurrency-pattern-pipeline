package main

import (
	"log"
	"time"

	dumyfile "github.com/faniafi/golang-concurrency-pattern-pipeline/dumyFile"
)

func main() {
	log.Println("Start")
	start := time.Now()

	dumyfile.GenerateFile()

	duration := time.Since(start)
	log.Printf("done id %.3f seconds\n", duration.Seconds())
}
