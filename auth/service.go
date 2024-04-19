package auth

import (
	"crypto/ecdsa"
	"database/sql"
	"errors"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/maxmx03/careconnect-backend/user"
)

type jwtCustomClaims struct {
	Email string `json:"Email"`
	Type  string `json:"Type"`
	jwt.RegisteredClaims
}

type AuthService struct{}

func (s *AuthService) Login(auth *AuthModel, db *sql.DB) (string, error) {
	query := "SELECT email, password, type FROM user where email = ?"
	user := &UserModel{}

	if err := db.QueryRow(query, auth.Email).Scan(&user.Email, &user.Password, &user.Type); err != nil {
		return "", err
	}

	if auth.Email == user.Email && auth.Password == user.Password {
		claims := &jwtCustomClaims{
			user.Email,
			user.Type,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		key, err := loadECDSAKey("ec256-private.pem")
		if err != nil {
			return "", err
		}
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		t, err := token.SignedString(key)

		if err != nil {
			return "", err
		}

		return t, nil
	}

	return "", errors.New("Unauthorized")
}

func loadECDSAKey(filePath string) (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseECPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return key, nil
}
