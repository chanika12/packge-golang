package packge_golang

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// ToUpperCase แปลงข้อความให้เป็นตัวพิมพ์ใหญ่
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}
func TestCha(s string) string {
	return fmt.Sprintf("Hello, %s!\n", s)
}

// GenerateJWT สร้าง JWT token จากข้อมูลและ secret key
func GenerateJWT(username, secret string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // หมดอายุใน 1 ชั่วโมง
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// VerifyJWT ตรวจสอบ JWT token ด้วย secret key
func VerifyJWT(tokenString, secret string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
