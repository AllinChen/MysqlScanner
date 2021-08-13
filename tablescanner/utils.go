package tablescanner

import (
	"encoding/json"
	"strconv"
)

// TransValue 将接口值的类型转换为string
func TransValue(value interface{}) string {
	var key string
	switch value := value.(type) {
	case float64:
		key = strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int:
		key = strconv.Itoa(value)
	case uint:
		key = strconv.Itoa(int(value))
	case int8:
		key = strconv.Itoa(int(value))
	case uint8:
		key = strconv.Itoa(int(value))
	case int16:
		key = strconv.Itoa(int(value))
	case uint16:
		key = strconv.Itoa(int(value))
	case int32:
		key = strconv.Itoa(int(value))
	case uint32:
		key = strconv.Itoa(int(value))
	case int64:
		key = strconv.FormatInt(value, 10)
	case uint64:
		key = strconv.FormatUint(value, 10)
	case string:
		key = value
	case []byte:
		key = string(value)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
