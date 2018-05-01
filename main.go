package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

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
	var myFlags arrayFlags

	out := flag.String("o", "", "The output destination of your zipped folder")
	flag.Var(&myFlags, "f", "The file you want to zip")

	flag.Parse()

	err := zipFiles(*out, myFlags)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Zipped File: " + *out)
}
