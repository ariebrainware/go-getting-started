package utils

import "fmt"

func Calculate(a, b int) (result int) {
	return a + b
}

func PingServer(ip string) (msg string, err error) {
	if ip != "" {
		return "success", nil
	}
	return "", fmt.Errorf("invalid ip address")
}
