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
	// menggunakan fan-out pattern untuk membaca dari data yang sama (dalam hal ini channel chanFileContent)
	chanFileSum1 := concurrency.GetSum(chanFileContent)
	chanFileSum2 := concurrency.GetSum(chanFileContent)
	chanFileSum3 := concurrency.GetSum(chanFileContent)
	chanFileSum4 := concurrency.GetSum(chanFileContent)
	chanFileSum := concurrency.MergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3, chanFileSum4)

	// pipeline 3 := rename file
	chanFileRename1 := concurrency.RenameFile(chanFileSum)
	chanFileRename2 := concurrency.RenameFile(chanFileSum)
	chanFileRename3 := concurrency.RenameFile(chanFileSum)
	chanFileRename4 := concurrency.RenameFile(chanFileSum)

	// menggunakan fan-in pattern untuk menggabungkan hasil dari beberapa task (dari task RenameFile) menjadi 1 result
	chanFileRename := concurrency.MergeChanFileInfo(chanFileRename1, chanFileRename2, chanFileRename3, chanFileRename4)

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
