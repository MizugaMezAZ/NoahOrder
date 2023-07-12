package service

import "gorder/api/repository"

type IAuthService interface {
}

type authService struct {
	ar repository.IAuthRepository
}

func NewAuthService(ar repository.IAuthRepository) IAuthService {
	return &authService{
		ar: ar,
	}
}

// ----------------------------------
