package format

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func testFile(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}

	var feature Feature

	json.Unmarshal(fileBytes, &feature)
	if err != nil {
		t.Error(err)
	}

	b, err := FormatFeature(&feature)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(b, fileBytes) != 0 {
		t.Errorf("Bytes did not match:\n%s", b)
	}
}

func TestFormatNewYork(t *testing.T) {
	testFile(t, "fixtures/85977539.geojson")
}

func TestFormatUnitedKingdom(t *testing.T) {
	testFile(t, "fixtures/85633159.geojson")
}
