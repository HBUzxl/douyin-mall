package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 签发token
func SignToken(userUuid string, privateKeyString string, tokenExpire int, refreshTokenExpire int) (string, string, error) {
	tokenExpireAt := time.Now().Add(time.Duration(tokenExpire) * time.Hour)
	refreshTokenExpireAt := time.Now().Add(time.Duration(refreshTokenExpire) * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  tokenExpireAt.Unix(),
			"rt":   false,
		},
	)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  refreshTokenExpireAt.Unix(),
			"rt":   true, // 是否是refreshToken
		},
	)

	privateKeyBytes, err := String2Hex(privateKeyString)
	if err != nil {
		return "", "", err
	}

	privateKey := ed25519.PrivateKey(privateKeyBytes)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}

	refreshtokenString, err := refreshToken.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshtokenString, nil
}

func String2Hex(s string) ([]byte, error) {
	if s == "" {
		return nil, errors.New("empty secret")
	}
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
