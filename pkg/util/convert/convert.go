package convert

import (
	"strconv"
)

func ToInt(v interface{}) int {
	var result int
	if v == nil {
		return result
	}
	switch v := v.(type) {
	case string:
		result, _ = strconv.Atoi(v)
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
	case []byte:
		result = ToInt(string(v))
	case bool:
		if v {
			result = 1
		}
	}
	return result
}
