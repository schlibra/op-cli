package utils

import (
	"fmt"
)

func FormatFileSize(bytes int64) string {
	if bytes < 0 {
		return "0 B"
	}
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	var i int
	val := float64(bytes)

	for val >= 1024 && i < len(units)-1 {
		val /= 1024
		i++
	}

	// 保留 2 位小数
	return fmt.Sprintf("%.2f %s", val, units[i])
}
