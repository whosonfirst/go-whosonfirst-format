package encoding

import (
	"fmt"
	"strings"

	"golang.org/x/text/unicode/norm"
)

type String struct {
	Value            string
	UnicodeNormalize string
}

func (s String) MarshalJSON() ([]byte, error) {
	str := s.Value

	switch strings.ToUpper(s.UnicodeNormalize) {
	case "NFD":
		str = norm.NFD.String(str)
	case "NFC":
		str = norm.NFC.String(str)
	case "NFKD":
		str = norm.NFKD.String(str)
	case "NFKC":
		str = norm.NFKC.String(str)
	}

	return []byte(fmt.Sprintf(`"%s"`, str)), nil
}
