# zip

Zip is a command line app to for zipping up files. 

## Installation

1. Clone this repo into `$GOPATH/src/github.com/ConradPacesa`. 

2. cd into the directory 

3. Build it with `go build -o $GOPATH/src/github.com/ConradPacesa/zip/bin/zip.exe main.go`

4. Add `$GOPATH/src/github.com/ConradPacesa/zip/bin/` to your system path variable. 

You're ready to go!

## Usage 

Usage is as simple as:

`zip -o <output-file> -f <files or folder to zip>`

### Example:

`zip -o data.zip -f data.csv work.csv info.csv`