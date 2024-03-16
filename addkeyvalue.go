package redisclone

import (
	"errors"
	"fmt"
)

var inMemoryData map[string]string

func AddKeyValue(key string, val string) {
	inMemoryData[key] = val
}

func GetValue(key string) (string, error) {
	if value, ok := inMemoryData[key]; ok {
		return value, nil
	}

	errMsg := fmt.Sprintf("Key: [%s] does not exist.", key)

	return "", errors.New(errMsg)
}
