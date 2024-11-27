package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	// Parse the string to uint64 first
	val, err := strconv.ParseUint(s, 10, 0) // Base 10, bit size 0 for uint
	if err != nil {
		return 0, err // Return 0 and the error if parsing fails
	}

	// Check for overflow if converting to uint
	if val > uint64(^uint(0)) {
		return 0, fmt.Errorf("value %d is too large for uint", val)
	}

	return uint(val), nil // Convert uint64 to uint and return
}

func Struct2Map(srcData any) (map[string]interface{}, error) {
	var trgData map[string]interface{}
	bytes, err := json.Marshal(srcData)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON into the map
	err = json.Unmarshal(bytes, &trgData)
	if err != nil {
		return nil, err
	}
	return trgData, nil
}

func GetStructProperty(srcData any, propertyName string) (*string, *bool, error) {
	mapData, err := Struct2Map(srcData)
	if err != nil {
		return nil, nil, err
	}
	value, exists := mapData[propertyName]
	if !exists {
		return nil, &exists, nil
	}
	str, ok := value.(string)
	if !ok {
		return nil, &exists, nil
	}
	return &str, &exists, nil
}
