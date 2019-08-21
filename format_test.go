package format

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
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

	formattedJSON := string(b)
	originalJSON := string(fileBytes)

	if strings.Compare(formattedJSON, originalJSON) != 0 {
		t.Errorf("Did not match:\n%v", diff.LineDiff(formattedJSON, originalJSON))
	}
}

func TestFormatNewYork(t *testing.T) {
	testFile(t, "fixtures/85977539.geojson")
}

func TestFormatUnitedKingdom(t *testing.T) {
	testFile(t, "fixtures/85633159.geojson")
}
