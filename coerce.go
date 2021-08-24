package format

import (
	"fmt"
	"strings"

	encoding "github.com/whosonfirst/go-whosonfirst-format/encoding"
)

func encode(path string, v interface{}) interface{} {
	var _type = fmt.Sprintf("%T", v)

	switch {
	case path == "$.properties.lbl:max_zoom" && _type == "float64":
		return encoding.Float{v.(float64), 1, 8}
	case strings.HasPrefix(path, "$.properties.name:deu_x_preferred") && _type == "string":
		return encoding.String{v.(string), ""}
	}

	return v
}

// type conversions
func Coerce(value interface{}, path string) interface{} {

	// collections
	switch value.(type) {

	// JSON object (string keys)
	case map[string]interface{}:
		var values = value.(map[string]interface{})
		for k, v := range values {
			values[k] = Coerce(v, fmt.Sprintf("%s.%s", path, k))
		}
		return values

	// JSON array (numeric keys)
	case []interface{}:
		var values = value.([]interface{})
		for k, v := range values {
			values[k] = Coerce(v, fmt.Sprintf("%s[%d]", path, k))
		}
		return values

	// scalar values
	default:
		return encode(path, value)

	}
}
