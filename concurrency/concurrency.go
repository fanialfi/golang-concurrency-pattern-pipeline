package concurrency

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/faniafi/golang-concurrency-pattern-pipeline/lib"
)

type FileInfo struct {
	FilePath  string // file location
	Content   []byte // content file
	Sum       string // md5 sum of content
	IsRenamed bool   // mengindikasikan apakah file yang terkait sudah direname
}

func ReadFile() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(lib.TempPath, func(path string, info fs.FileInfo, err error) error {
			// jika terjadi error maka langsung return
			if err != nil {
				return err
			}

			// jika ini adalah sebuah sub-directory maka langsung return
			if info.IsDir() {
				return nil
			}

			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil
		})

		if err != nil {
			log.Println("ERROR :", err.Error())
		}

		close(chanOut)
	}()

	return chanOut
}

func GetSum(chanInput <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		// listen setiap ada penerimaan data dari channel chanInput
		// dengan cara di-loop & loop ini akan otomatis berhenti ketika tidak ada data lagi
		// yang masuk ke channel chanInput / chanInput sudah di close
		for fileInfo := range chanInput {
			// dari struct FileInfo, tambahkan value untuk field Sum
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}

		// ketika sudah tidak ada penerimaan data dari channel chanInput, / chanInput sudah diclose
		// maka channel chanOut di close
		close(chanOut)
	}()

	return chanOut
}

func RenameFile(chanFileInfo <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanFileInfo {
			newPath := filepath.Join(lib.TempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}

		close(chanOut)
	}()

	return chanOut
}
