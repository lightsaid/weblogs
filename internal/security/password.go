package security

import "golang.org/x/crypto/bcrypt"

// Hash 生成hash密码
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword 验证hash密码
func VerifyPassword(hashedPassword, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
}
