package service

type Repository interface {
}

type PasswordUtils interface {
	GeneratePassword(password string) (string, error)
	ComparePassword(password, hash string) (bool, error)
}

type JWTUtils interface {
	GenerateAccessToken(userId int) (string, int64, error)
	GenerateRefreshToken(userId int) (string, int64, error)
	VerifyToken(token string) (int, int64, error)
}
