package encoding

import (
	"bytes"
	"strconv"
)

type Float struct {
	Value float64
	Min   int
	Max   int
}

func (f Float) MarshalJSON() ([]byte, error) {

	// format to maximum precision (use -1 for unlimited)
	enc := []byte(strconv.FormatFloat(float64(f.Value), 'f', f.Max, 64))

	// remove trailing 0s while maintaining $Min precision
	if f.Min > -1 {
		pos := bytes.IndexByte(enc, '.')
		if f.Min == 0 {
			return enc[:pos], nil
		}
		return enc[:pos+f.Min+1], nil
	}

	return enc, nil
}
