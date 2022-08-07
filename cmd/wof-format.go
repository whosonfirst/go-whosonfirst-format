package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	format "github.com/whosonfirst/go-whosonfirst-format"
)

func main() {
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
		fmt.Fprintf(os.Stderr, "You cannot combine the -output and -overwrite flags\n")
		os.Exit(1)
	}

	if *checkFlag && *outputFlag != "" {
		fmt.Fprintf(os.Stderr, "You cannot combine the -check and -output flags\n")
		os.Exit(1)
	}

	if *checkFlag && *overwriteFlag {
		fmt.Fprintf(os.Stderr, "You cannot combine the -check and -overwrite flags\n")
		os.Exit(1)
	}

	path, reader, err := inputReader()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}

	inputBytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}

	outputBytes, err := format.FormatBytes(inputBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}

	if *checkFlag {
		if bytes.Equal(inputBytes, outputBytes) {
			os.Exit(0)
			return
		}

		if path != "" {
			fmt.Fprintf(os.Stderr, "%s is not valid\n", path)
		}

		os.Exit(1)
	}

	reader.Close()

	var output io.WriteCloser = os.Stdout

	if *outputFlag != "" {
		output, err = os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open %s for writing\n", *outputFlag)
			os.Exit(1)
		}
	}

	if *overwriteFlag && path != "" {
		output, err = os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open %s for writing\n", *outputFlag)
			os.Exit(1)
		}
	}

	buffer := bufio.NewWriter(output)

	_, err = buffer.Write(outputBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write\n")
		os.Exit(1)
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
