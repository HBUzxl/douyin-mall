package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"k8s.io/klog/v2"
)

// 签发token
func SignToken(userUuid string, privateKeyString string, tokenExpire int64, refreshTokenExpire int64) (string, string, error) {
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

// 验证Token
func VertifyToken(tokenString, publicKeyHexString string) (string, bool, error) {
	publicKeyBytes, err := String2Hex(publicKeyHexString)
	if err != nil {
		return "", false, err
	}

	publicKey := ed25519.PublicKey(publicKeyBytes)

	// 解析验证Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return "", false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, err
	}
	expireAt := claims["exp"].(float64)
	if time.Now().Unix() > int64(expireAt) {
		klog.Errorf("token expired")
		return "", false, errors.New("token expired")
	}

	isRefreshToken := claims["rt"].(bool)

	return claims["uuid"].(string), isRefreshToken, nil
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
