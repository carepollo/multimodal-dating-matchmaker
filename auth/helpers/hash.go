package helpers

import (
	"time"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// hash and salt plain string password, returns the hashed password
func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// check if plain string password is the equivalent of a given hash
func CheckPassword(plain string, hashed string) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plain))
	return err == nil
}

// generar un jwt con el user id almacenado dentro
func GenerateToken(userId string, key string) (string, error) {
	claims := models.LoggedInClient{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signed, nil
}
