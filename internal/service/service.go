package service

type Service struct {
	repository Repository
	passUtils  PasswordUtils
	jwtUtils   JWTUtils
}

func New(repo Repository, passUtils PasswordUtils, jwtUtils JWTUtils) *Service {
	return &Service{
		repository: repo,
		passUtils:  passUtils,
		jwtUtils:   jwtUtils,
	}
}
