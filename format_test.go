package format

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/paulmach/orb/geojson"
)

const fixturesPath = "fixtures"

func testFile(t *testing.T, inputPath string, expectedOutputPath string) {
	inputBytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Error(err)
	}

	expectedBytes, err := ioutil.ReadFile(expectedOutputPath)
	if err != nil {
		t.Error(err)
	}

	feature, err := geojson.UnmarshalFeature(inputBytes)

	if err != nil {
		t.Error(err)
	}

	b, err := FormatFeature(feature)
	if err != nil {
		t.Error(err)
	}

	formattedJSON := string(b)
	expectedJSON := string(expectedBytes)

	if strings.Compare(expectedJSON, formattedJSON) != 0 {
		d := diff.LineDiff(expectedJSON, formattedJSON)
		t.Errorf("%s and %s did not match:\n%v", inputPath, expectedOutputPath, d)
	} else {
		t.Logf("%s and %s matched", inputPath, expectedOutputPath)
	}
}

func TestFormat(t *testing.T) {
	files, err := os.ReadDir(fixturesPath)
	if err != nil {
		t.Error(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		n := file.Name()

		if !strings.HasSuffix(n, ".geojson") {
			continue
		}

		if strings.HasSuffix(n, ".expected.geojson") {
			continue
		}

		inputPath := path.Join(fixturesPath, n)
		expectedPath := path.Join(fixturesPath, strings.Replace(n, ".geojson", ".expected.geojson", 1))

		testFile(t, inputPath, expectedPath)
	}

}
