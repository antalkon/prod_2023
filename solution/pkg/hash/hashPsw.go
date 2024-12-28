package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Генерация хэша пароля с использованием bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword принимает пароль и хэш, и проверяет их соответствие
func CheckPassword(password, hash string) error {
	// Сравнивает введённый пароль с хэшированным паролем
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
