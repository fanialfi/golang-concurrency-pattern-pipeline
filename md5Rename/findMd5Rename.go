package md5rename

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/faniafi/golang-concurrency-pattern-pipeline/lib"
)

func Proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(lib.TempPath, func(path string, info fs.FileInfo, err error) error {
		// jika terjadi error di argument err maka langsung return
		if err != nil {
			return err
		}

		// jika itu adalah sebuah sub directory / directory maka juga langsung return
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// read file
		buffer, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// cari md5 sum-nya
		md5Sum := fmt.Sprintf("%x", md5.Sum(buffer))

		// rename file
		destinationPath := filepath.Join(lib.TempPath, fmt.Sprintf("file-%s.txt", md5Sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR :", err.Error())
	}

	log.Printf("%d/%d file reanamed", counterRenamed, counterTotal)
}
