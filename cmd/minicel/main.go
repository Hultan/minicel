package main

import (
	"fmt"
	"os"

	"github.com/hultan/minicel/internal/minicel"
)

type ExitCode int

const (
	ExitCodeIllegalUsage ExitCode = 1
	ExitCodeFileNotExist
	ExitCodeParseError
)

type Arguments int

const (
	_ Arguments = iota
	ArgumentFileName
)

func main() {
	if len(os.Args) != 2 {
		usage(os.Stderr)
		_, _ = fmt.Fprintf(os.Stderr, "ERROR : Input file is not provided.\n")
		os.Exit(int(ExitCodeIllegalUsage))
	}

	fileName := os.Args[ArgumentFileName]
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR : Input file '%s' does not exist.\n", fileName)
		os.Exit(int(ExitCodeFileNotExist))
	}

	mc := minicel.NewMinicel()
	records, err := mc.Evaluate(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR : Input file '%s' has invalid syntax.\n", fileName)
		os.Exit(int(ExitCodeParseError))
	}

	// Print result
	fmt.Fprintf(os.Stdout, "\n")
	for _, row := range records {
		for i, cell := range row {
			if i > 0 {
				fmt.Fprintf(os.Stdout, ",")
			}
			fmt.Fprintf(os.Stdout, "%s", cell.Value())
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}

func usage(stderr *os.File) {
	_, _ = fmt.Fprintf(stderr, "Usage : minicel [input file]")
}
