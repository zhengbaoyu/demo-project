package jwt

import "github.com/golang-jwt/jwt/v4"

func GetToken(iat int64, secretKey string, payloads map[string]any, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["expTime"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
