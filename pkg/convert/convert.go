package convert

import (
	"strconv"
	"strings"
)

func ToInt(v interface{}) int {
	var result int
	switch v := v.(type) {
	case int:
		result = v
	case int32:
		result = int(v)
	case int64:
		result = int(v)
	case float32:
		result = int(v)
	case float64:
		result = int(v)
	case string:
		result, _ = strconv.Atoi(strings.TrimSpace(v))
	case []byte:
		result = ToInt(string(v))
	}
	return result
}
