# go-whosonfirst-format

Standardised GeoJSON formatting for Whos On First files.

Usable as both a library and a binary.

## Library usage

```golang
func main() {
  inputBytes, err := ioutil.ReadFile(inputPath)
  if err != nil {
    panic(err)
  }

  var feature format.Feature

  json.Unmarshal(inputBytes, &feature)
  if err != nil {
    panic(err)
  }

  outputBytes, err := format.FormatFeature(&feature)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s", outputBytes)
}
```

## Tools

```shell
$> make build
go build -mod vendor -ldflags="-s -w" -o bin/wof-format ./cmd/wof-format/main.go
```

### wof-format

```shell
$> ./bin/wof-format -h
Usage: ./bin/wof-format [-check] [input]

Either provide a WOF feature on stdin:
cat input.geojson | ./bin/wof-format > output.geojson

Or provide a path to a WOF feature as the first argument:
./bin/wof-format input.geojson > output.geojson

Optional arguments:

  -check
    	exits silently with a non-zero status code if the file is not correctly formatted
  -output string
    	write the output to a file instead of stdout
  -overwrite
    	overwrite the input file with the formatted output
``	

For example:

```
$> cat input.geojson | ./build/wof-format > output.geojson
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-format-wasm