package auth

import (
	"crypto/rand"
	"crypto/sha3"
)

type user struct {
	username string
	password string
	salt     string
}

func (u *user) generatePasswordHash(password string) error {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	hash := makeHash(password, salt)

	u.password = hash
	u.salt = string(salt)
	return nil
}

func (u *user) verifyPassword(password string) bool {
	hash := makeHash(password, []byte(u.salt))

	return hash == u.password

}

func makeHash(password string, salt []byte) string {

	algo := sha3.New256()
	saltedPass := append([]byte(password), salt...)
	hash := algo.Sum(saltedPass)
	return string(hash)

}
