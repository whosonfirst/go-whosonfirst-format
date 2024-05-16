package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	format "github.com/whosonfirst/go-whosonfirst-format"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	flag.CommandLine.SetOutput(os.Stderr)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-check] [input]\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Either provide a WOF feature on stdin:\ncat input.geojson | %s > output.geojson\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Or provide a path to a WOF feature as the first argument:\n%s input.geojson > output.geojson\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Optional arguments:\n\n")
		flag.PrintDefaults()
	}

	checkFlag := flag.Bool("check", false, "exits silently with a non-zero status code if the file is not correctly formatted")
	outputFlag := flag.String("output", "", "write the output to a file instead of stdout")
	overwriteFlag := flag.Bool("overwrite", false, "overwrite the input file with the formatted output")
	flag.Parse()

	if *outputFlag != "" && *overwriteFlag {
		l.Fatalf("You cannot combine the -output and -overwrite flags\n")
		return
	}

	if *checkFlag && *outputFlag != "" {
		l.Fatalf("You cannot combine the -check and -output flags\n")
		return
	}

	if *checkFlag && *overwriteFlag {
		l.Fatalf("You cannot combine the -check and -overwrite flags\n")
		return
	}

	path, reader, err := inputReader()
	if err != nil {
		l.Fatalf("Error reading: %v\n", err)
		return
	}

	inputBytes, err := io.ReadAll(reader)
	if err != nil {
		l.Fatalf("%v\n", err)
		return
	}

	outputBytes, err := format.FormatBytes(inputBytes)
	if err != nil {
		l.Fatalf("%v\n", err)
		return
	}

	if *checkFlag {
		if bytes.Equal(inputBytes, outputBytes) {
			os.Exit(0)
			return
		}

		if path != "" {
			l.Printf("%s is not valid\n", path)
		}

		os.Exit(1)
	}

	reader.Close()

	var output io.WriteCloser = os.Stdout

	if *outputFlag != "" {
		output, err = os.Create(*outputFlag)
		if err != nil {
			l.Fatalf("Unable to open %s for writing\n", *outputFlag)
			return
		}
	}

	if *overwriteFlag && path != "" {
		output, err = os.Create(path)
		if err != nil {
			l.Fatalf("Unable to open %s for writing\n", *outputFlag)
			return
		}
	}

	buffer := bufio.NewWriter(output)

	_, err = buffer.Write(outputBytes)
	if err != nil {
		log.Fatalf("Unable to write: %v\n", err)
		return
	}

	buffer.Flush()
	output.Close()
}

func inputReader() (string, io.ReadCloser, error) {
	if flag.NArg() == 0 {
		info, err := os.Stdin.Stat()
		if err != nil {
			flag.Usage()
			os.Exit(1)
		}

		if (info.Mode() & os.ModeCharDevice) != 0 {
			flag.Usage()
			os.Exit(1)
		}

		return "", os.Stdin, nil
	}

	path := flag.Arg(0)
	f, err := os.Open(path)
	if err != nil {
		return path, nil, fmt.Errorf("unable to open file: %w", err)
	}

	return path, f, nil
}
