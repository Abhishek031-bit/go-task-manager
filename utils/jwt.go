package utils

import (
	"errors"
	"fmt"
	"task-manager/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtBlacklist = make(map[string]bool)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", errors.New("invalid token expiration")
	}

	if time.Now().Unix() > int64(exp) {
		return "", errors.New("token has expired")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("invalid user ID in token")
	}

	return GenerateToken(uint(userID))
}

func BlacklistToken(token string) {
	jwtBlacklist[token] = true
}

func IsTokenBlacklisted(token string) bool {
	return jwtBlacklist[token]
}
