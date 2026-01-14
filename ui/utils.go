package ui

import (
	"fmt"
	"strings"
)

// FormatNumber 格式化数字，去除不必要的尾零
func FormatNumber(val float64) string {
	s := fmt.Sprintf("%.10f", val)
	for len(s) > 1 && s[len(s)-1] == '0' && s[len(s)-2] != '.' {
		s = s[:len(s)-1]
	}
	if s[len(s)-1] == '.' {
		s = s + "0"
	}
	return s
}

// FormatTokens 格式化tokens数量，大数字使用K/M单位，去除尾部多余的零
func FormatTokens(val float64) string {
	var num float64
	var suffix string

	if val >= 1000000 {
		num = val / 1000000
		suffix = "M"
	} else if val >= 1000 {
		num = val / 1000
		suffix = "K"
	} else {
		num = val
		suffix = ""
	}

	// 格式化数字，去除尾部多余的零
	formatted := fmt.Sprintf("%.2f", num)
	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")

	return formatted + suffix
}
