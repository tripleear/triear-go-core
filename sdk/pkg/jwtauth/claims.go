package jwtauth

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

// MapClaims represents a jwt.MapClaims
type MapClaims jwt.MapClaims

// Exp returns value of exp
func (m MapClaims) Exp() (int64, error) {
	return m.Int64("exp")
}

// OrigIat returns value of orig_iat
func (m MapClaims) OrigIat() (int64, error) {
	return m.Int64("orig_iat")
}

// Identity returns value of identity
func (m MapClaims) Identity() (int64, error) {
	return m.Int64("identity")
}

// Int64 try to convert to int64 by key
func (m MapClaims) Int64(key string) (int64, error) {
	value := m[key]
	if value == nil {
		return 0, fmt.Errorf("invalid key '%v'", key)
	}

	switch value.(type) {
	case json.Number:
		return value.(json.Number).Int64()
	case float64:
		return int64(value.(float64)), nil
	case string:
		return strconv.ParseInt(value.(string), 10, 0)
	default:
		return 0, fmt.Errorf("invalid value '%v' type '%T'", value, value)
	}
}

// String try to convert to string by key
func (m MapClaims) String(key string) string {
	value := m[key]
	if value == nil {
		return ""
	}

	switch value.(type) {
	case json.Number:
		return value.(json.Number).String()
	case float64:
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	case string:
		return value.(string)
	default:
		fmt.Printf("maptoclaims key to string invalid value '%v' type '%T' key %v", value, value, key)
		return ""
	}
}

// Int try to convert to int by key
func (m MapClaims) Int(key string) (int, error) {
	value := m[key]
	if value == nil {
		return 0, fmt.Errorf("invalid key '%v'", key)
	}

	switch value.(type) {
	case json.Number:
		intVal, err := value.(json.Number).Int64()
		if err != nil {
			return 0, err
		}
		return int(intVal), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		intVal, err := strconv.ParseInt(value.(string), 10, 0)
		if err != nil {
			return 0, err
		}
		return int(intVal), nil
	default:
		return 0, fmt.Errorf("invalid value '%v' type '%T'", value, value)
	}
}

// UInt64 try to convert to uint64 by key
func (m MapClaims) UInt64(key string) (uint64, error) {
	value := m[key]
	if value == nil {
		return 0, fmt.Errorf("invalid key '%v'", key)
	}

	switch value.(type) {
	case json.Number:
		intVal, err := value.(json.Number).Int64()
		if err != nil {
			return 0, err
		}
		return uint64(intVal), nil
	case float64:
		return uint64(value.(float64)), nil
	case string:
		intVal, err := strconv.ParseInt(value.(string), 10, 0)
		if err != nil {
			return 0, err
		}
		return uint64(intVal), nil
	default:
		return 0, fmt.Errorf("invalid value '%v' type '%T'", value, value)
	}
}

// Int64Array try to convert to []int64 by key
func (m MapClaims) Int64Array(key string) ([]int64, error) {
	raw := m[key]
	if raw == nil {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	raw, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	rawList, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("error format, key '%v'", key)
	}

	var vals []int64
	for _, r := range rawList {
		switch r.(type) {
		case json.Number:
			intVal, err := r.(json.Number).Int64()
			if err != nil {
				return nil, err
			}
			vals = append(vals, intVal)
		case float64:
			vals = append(vals, int64(r.(float64)))
		case string:
			intVal, err := strconv.ParseInt(r.(string), 10, 0)
			if err != nil {
				return nil, err
			}
			vals = append(vals, intVal)
		default:
			return nil, fmt.Errorf("invalid value '%v' type '%T'", r, r)
		}
	}

	return vals, nil
}

// StringArray try to convert to []string by key
func (m MapClaims) StringArray(key string) []string {
	raw := m[key]
	if raw == nil {
		return nil
	}

	raw, ok := m[key]
	if !ok {
		return nil
	}

	rawList, ok := raw.([]interface{})
	if !ok {
		return nil
	}

	var vals []string
	for _, r := range rawList {
		switch r.(type) {
		case json.Number:
			vals = append(vals, r.(json.Number).String())
		case float64:
			vals = append(vals, strconv.FormatFloat(r.(float64), 'g', -1, 64))
		case string:
			vals = append(vals, r.(string))
		default:
			fmt.Printf("maptoclaims key to string invalid value '%v' type '%T' key %v", r, r, key)
			vals = append(vals, "")
		}
	}
	return vals
}

// IntArray try to convert to []int by key
func (m MapClaims) IntArray(key string) ([]int, error) {
	raw := m[key]
	if raw == nil {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	raw, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	rawList, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("error format, key '%v'", key)
	}

	var vals []int
	for _, r := range rawList {
		switch r.(type) {
		case json.Number:
			intVal, err := r.(json.Number).Int64()
			if err != nil {
				return nil, err
			}
			vals = append(vals, int(intVal))
		case float64:
			vals = append(vals, int(r.(float64)))
		case string:
			intVal, err := strconv.Atoi(r.(string))
			if err != nil {
				return nil, err
			}
			vals = append(vals, intVal)
		default:
			return nil, fmt.Errorf("invalid value '%v' type '%T'", r, r)
		}
	}

	return vals, nil
}

// UInt64Array try to convert to []uint64 by key
func (m MapClaims) UInt64Array(key string) ([]uint64, error) {
	raw := m[key]
	if raw == nil {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	raw, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("invalid key '%v'", key)
	}

	rawList, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("error format, key '%v'", key)
	}

	var vals []uint64
	for _, r := range rawList {
		switch r.(type) {
		case json.Number:
			intVal, err := r.(json.Number).Int64()
			if err != nil {
				return nil, err
			}
			vals = append(vals, uint64(intVal))
		case float64:
			vals = append(vals, uint64(r.(float64)))
		case string:
			intVal, err := strconv.Atoi(r.(string))
			if err != nil {
				return nil, err
			}
			vals = append(vals, uint64(intVal))
		default:
			return nil, fmt.Errorf("invalid value '%v' type '%T'", r, r)
		}
	}

	return vals, nil
}
