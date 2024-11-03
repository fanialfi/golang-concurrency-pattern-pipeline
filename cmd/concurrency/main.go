package main

import (
	"log"
	"time"

	"github.com/faniafi/golang-concurrency-pattern-pipeline/concurrency"
)

func main() {
	log.Println("Start")
	start := time.Now()

	// pipeline 1 := loop all file and read it
	chanFileContent := concurrency.ReadFile()

	// pipeline 2 := calculate md5 sum
	chanFileSum := concurrency.GetSum(chanFileContent)

	// pipeline 3 := rename file
	chanFileRename := concurrency.RenameFile(chanFileSum)

	// pipeline 4 := print output
	counterRenamed := 0
	counterTotal := 0

	for fileInfo := range chanFileRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d file renamed\n", counterRenamed, counterTotal)
	duration := time.Since(start)
	log.Printf("done id %.3f seconds\n", duration.Seconds())
}
