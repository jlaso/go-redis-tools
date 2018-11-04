package common

import "fmt"

func GetServers() []string {
	var result []string
	cfg := GetConfiguration()
	result = append(result, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))

	return result
}
