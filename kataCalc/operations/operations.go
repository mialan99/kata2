package operations

import (
	"strconv"
	"strings"
)

func Sum(a string, b string) string {
	a = strings.Trim(a, `"`)
	b = strings.Trim(b, `"`)
	return `"` + a + b + `"`
}

func Sub(a, b string) string {
	a = strings.Trim(a, `"`)
	b = strings.Trim(b, `"`)
	if len(b) > len(a) {
		return `"` + a + `"`
	}
	if a == b {
		return ""
	}
	for i := 0; i < len(a); i++ {
		j := i + len(b)
		if j > len(a) {
			return `"` + a + `"`
		}
		if string(a[i:j]) == b {
			a = string(a[:i]) + string(a[j:])
			return `"` + a + `"`
		}
	}
	return `"` + a + `"`
}

func Multiply(a, b string) string {
	a = strings.Trim(a, `"`)
	bNum, _ := strconv.Atoi(b)
	result := strings.Repeat(a, bNum)
	if len(result) > 40 {
		result = result[0:40] + "..."
	}
	return `"` + result + `"`
}

func Div(a, b string) string {
	a = strings.Trim(a, `"`)
	bNum, _ := strconv.Atoi(b)
	if bNum > len(a) {
		panic("Деление невозможно. Символов меньше чем значение знаменателя")
	}
	return string(a[:len(a)/bNum])
}
