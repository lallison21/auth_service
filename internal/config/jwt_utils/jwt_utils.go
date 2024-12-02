package jwt_utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTConfig struct {
	SecretKey       string        `env:"JWT_SECRET_KEY" required:"true" default:"superSecretKey"`
	AccessTokenExp  time.Duration `env:"JWT_ACCESS_TOKEN_EXP" required:"true" default:"5m"`
	RefreshTokenExp time.Duration `env:"JWT_REFRESH_TOKEN_EXP" required:"true" default:"1h"`
}

type Utils struct {
	cfg *JWTConfig
}

func New(cfg *JWTConfig) *Utils {
	return &Utils{
		cfg: cfg,
	}
}

func (u *Utils) GenerateAccessToken(userId int) (string, int64, error) {
	return u.generateToken(userId, u.cfg.SecretKey, u.cfg.AccessTokenExp)
}

func (u *Utils) GenerateRefreshToken(userId int) (string, int64, error) {
	return u.generateToken(userId, u.cfg.SecretKey, u.cfg.RefreshTokenExp)
}

func (u *Utils) VerifyToken(token string) (int, int64, error) {
	return u.verifyToken(token)
}

func (u *Utils) generateToken(userId int, secret string, duration time.Duration) (string, int64, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	exp := time.Now().Add(duration).Unix()
	claims["user_id"] = userId
	claims["exp"] = exp

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

func (u *Utils) verifyToken(tokenString string) (int, int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(u.cfg.SecretKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, 0, fmt.Errorf("invalid token")
	}

	exp := int64(claims["exp"].(float64))
	if time.Unix(exp, 0).Before(time.Now()) {
		return 0, 0, fmt.Errorf("token is expired")
	}

	userId := int(claims["user_id"].(float64))

	return userId, exp, nil
}
