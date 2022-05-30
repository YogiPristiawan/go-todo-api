package encrypt

import "golang.org/x/crypto/bcrypt"

type HashPassword struct {
}

func (h *HashPassword) ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h *HashPassword) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hashed)
}
