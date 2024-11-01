package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faniafi/golang-concurrency-pattern-pipeline/lib"
)

const (
	totalFile     = 3000
	contentLength = 5000
)

var tempPath = "/tmp/go-pipeline"

func main() {
	log.Println("start")
	start := time.Now()

	generateFile()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func generateFile() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := lib.RandomString(contentLength)

		err := os.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName, "with error :", err.Error())
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "file created")
		}
	}

	fmt.Printf("%d file created\n", totalFile)
}
