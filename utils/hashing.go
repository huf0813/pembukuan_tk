package utils

import "golang.org/x/crypto/bcrypt"

type Hashing struct{}

type HashingInterface interface {
	HashPass(pass string) ([]byte, error)
	ComparePass(hashedPass, pass string) error
}

func (h *Hashing) HashPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func (h *Hashing) ComparePass(hashedPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
