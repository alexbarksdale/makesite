package utils

import (
	"flag"
	"strings"
)

func Flags() string {
	var file string

	flag.StringVar(&file, "file", "", "Provide the path to the file to convert.")
	flag.Parse()

	if !strings.Contains(file, ".txt") {
		return "You must provide a valid .txt file. Use the flag '--file' for help."
	}

	return file
}
