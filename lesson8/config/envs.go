package config

import (
	"os"
	"strconv"
	"strings"
)

func EnvironsMap() map[string]string {
	items := make(map[string]string)
	for _, item := range os.Environ() {
		pair := strings.SplitN(item, "=", 2)
		key, val := pair[0], pair[1]
		items[key] = val
	}
	return items
}

func GetEnval(key string, defaultValue string) string {
	if enval, ok := os.LookupEnv(key); ok {
		return enval
	}
	return defaultValue
}

func GetEnvalInt64(key string, defaultValue int64) int64 {
	if enval, ok := os.LookupEnv(key); ok {
		envalInt64, err := strconv.ParseInt(enval, 10, 64)
		if err == nil {
			return envalInt64
		}
	}
	return defaultValue
}
