package utils

import "os"

func Contains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}

func Getenv(key, defaultval string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultval
}
