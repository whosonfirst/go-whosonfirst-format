package main

import (
	"bufio"
	"encoding/json"
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
			fmt.Printf("Unable to open stdin: %s", err)
			os.Exit(1)
			return
		}

		if (info.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("Usage: cat input.geojson | wof-format > output.geojson")
			os.Exit(1)
			return
		}

		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Printf("Unable to open file: %s", err)
			os.Exit(1)
			return
		}

		reader = bufio.NewReader(f)
	}

	inputBytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Unable to read: %s", err)
		os.Exit(1)
		return
	}

	var feature format.Feature
	err = json.Unmarshal(inputBytes, &feature)
	if err != nil {
		fmt.Printf("Invalid JSON: %s", err)
		os.Exit(1)
		return
	}

	outputBytes, err := format.FormatFeature(&feature)
	if err != nil {
		fmt.Printf("Failed to format feature: %s", err)
		os.Exit(1)
		return
	}

	fmt.Printf("%s", outputBytes)
}
