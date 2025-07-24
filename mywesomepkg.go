package packge_golang

import (
	"fmt"
	"strings"
)

// ToUpperCase แปลงข้อความให้เป็นตัวพิมพ์ใหญ่
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}
func TestCha(s string) string {
	return fmt.Sprintf("Hello, %s!\n", s)
}
