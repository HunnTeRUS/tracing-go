package main

import (
	"os"
	"strconv"
)

func GetEnvVar(key string, fallbackValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return fallbackValue
}

func GetBoolEnvVar(key string) bool {
	value := os.Getenv(key)
	boolVar, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolVar
}
