package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	format "github.com/whosonfirst/go-whosonfirst-format"
)

func main() {
	flag.Parse()

	var reader *bufio.Reader

	if flag.NArg() == 0 {
		info, err := os.Stdin.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open stdin: %s\n", err)
			os.Exit(1)
			return
		}

		if (info.Mode() & os.ModeCharDevice) != 0 {
			fmt.Fprintf(os.Stderr, "Usage: cat input.geojson | wof-format > output.geojson\n")
			os.Exit(1)
			return
		}

		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
			os.Exit(1)
			return
		}

		reader = bufio.NewReader(f)
	}

	inputBytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read: %s\n", err)
		os.Exit(1)
		return
	}

	outputBytes, err := format.FormatBytes(inputBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format feature: %s\n", err)
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stdout, "%s", outputBytes)
}
