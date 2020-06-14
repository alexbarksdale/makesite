package utils

import (
	"flag"
)

// TODO: Use the flag to find all `.txt` files in the given directory. Print them to `stdout`.

func Flags() string {
	var file string
	var dir string

	flag.StringVar(&file, "file", "", "Provide the path to the file to convert.")
	flag.StringVar(&dir, "dir", "", "Finds all '.txt' files in the given directory.")
	flag.Parse()

	return file
}
