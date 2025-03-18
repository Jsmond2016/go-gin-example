package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired = errors.New("令牌已过期")
	ErrTokenInvalid = errors.New("无效的令牌")
)

// Claims 自定义的 JWT 声明
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.JWT.ExpireDuration)

	claims := Claims{
		Username: EncodeMD5(username),
		Password: EncodeMD5(password),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    config.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWT.Secret)
}

// ParseToken 解析 JWT 令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return config.JWT.Secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}


// getUsernameFromClaims 从 claims 中提取 username
func GetUsernameFromClaims(c *gin.Context) (string, error) {
	claims, exist := c.Get("claims")
	if !exist {
		return "", fmt.Errorf("claims not found")
	}

	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid claims type")
	}

	username, usernameExists := claimsMap["username"]
	if !usernameExists {
		return "", fmt.Errorf("username not found in claims")
	}

	usernameStr, ok := username.(string)
	if !ok {
		return "", fmt.Errorf("invalid username type")
	}

	return usernameStr, nil
}