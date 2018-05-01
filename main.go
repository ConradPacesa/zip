package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ConradPacesa/zip/zip"
)

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

	if len(fileSlice) > 1 {
		err := zip.ZipFiles(*out, fileSlice)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := zip.ZipDir(*out, *file)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Zipped File: " + *out)
}
