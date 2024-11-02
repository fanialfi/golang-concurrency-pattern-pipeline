package main

import (
	"log"
	"time"

	md5rename "github.com/faniafi/golang-concurrency-pattern-pipeline/md5Rename"
)

func main() {
	log.Println("Start")
	start := time.Now()

	md5rename.Proceed()
	// dumyfile.GenerateFile()

	duration := time.Since(start)
	log.Printf("done id %.3f seconds\n", duration.Seconds())
}
