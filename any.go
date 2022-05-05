package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToString(val interface{}) string {
	switch x := val.(type) {
	case bool:
		if x {
			return "1"
		}
		return "0"
	case int:
		return fmt.Sprintf("%d", x)
	case int8:
		return fmt.Sprintf("%d", x)
	case int16:
		return fmt.Sprintf("%d", x)
	case int32:
		return fmt.Sprintf("%d", x)
	case int64:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case uint8:
		return fmt.Sprintf("%d", x)
	case uint16:
		return fmt.Sprintf("%d", x)
	case uint32:
		return fmt.Sprintf("%d", x)
	case uint64:
		return fmt.Sprintf("%d", x)
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case float64:
		if math.Floor(x) == x {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case string:
		return x
	case []byte:
		return string(x)
	case nil:
		return ""
	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return ""
		}

		return string(buf)
	}
}

func ToInt(val interface{}) int {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

func ToBool(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		return x
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return x != 0
	case string:
		switch strings.ToLower(x) {
		case "true", "1":
			return true
		case "false", "0":
			return false
		case "":
			return false
		default:
			return true
		}
	case []byte:
		switch string(bytes.ToLower(x)) {
		case "true", "1":
			return true
		case "false", "0":
			return false
		case "":
			return false
		default:
			return true
		}
	default:
		return val == nil
	}
}