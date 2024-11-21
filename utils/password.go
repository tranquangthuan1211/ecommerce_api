package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Sử dụng hàm GenerateFromPassword để mã hóa mật khẩu với chi phí 14 (cost)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Hàm xác thực mật khẩu
func CheckPassword(hashedPassword, password string) bool {
	// Sử dụng hàm CompareHashAndPassword để xác thực
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
