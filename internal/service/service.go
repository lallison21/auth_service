package service

type Service struct {
	repository Repository
}

func New(repo Repository) *Service {
	return &Service{
		repository: repo,
	}
}
