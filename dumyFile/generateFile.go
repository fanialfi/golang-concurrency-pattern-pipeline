package dumyfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/faniafi/golang-concurrency-pattern-pipeline/lib"
)

func GenerateFile() {
	os.RemoveAll(lib.TempPath)
	os.MkdirAll(lib.TempPath, os.ModePerm)

	for i := 0; i < lib.TotalFile; i++ {
		fileName := filepath.Join(lib.TempPath, fmt.Sprintf("file-%d.txt", i))
		content := lib.RandomString(lib.ContentLength)

		err := os.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName, "with error :", err.Error())
		}
	}

	fmt.Printf("%d file created\n", lib.TotalFile)
}
