package util

import (
	"golang.org/x/crypto/bcrypt"
)

// this function receives a string, turns it to bytes, hashes it and
// returns the hash result in string
func HashAndSalt(password string) string {
	bytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// check whether or not a plain string is equivalent to certain hash
func ComparePasswords(hashed string, plainPassword []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}
