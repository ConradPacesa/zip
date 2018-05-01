package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func zipFiles(filename string, files []string) error {
	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	for _, file := range files {
		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	out := flag.String("o", "", "The output destination of your zipped folder (required)")
	file := flag.String("f", "", "The file(s) you want to zip (required)")

	flag.Parse()

	if !strings.HasSuffix(*out, ".zip") {
		log.Fatal("You must specify a valid outfile with a .zip extension")
	}

	if *file == "" {
		log.Fatal("You must specify at least one file or folder to zip up.")
	}

	fileSlice := flag.Args()
	fileSlice = append(fileSlice, *file)

	err := zipFiles(*out, fileSlice)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Zipped File: " + *out)
}
