package usecase

import (
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/repository"
)

type usecase struct {
	repository repository.Repository
}

type Usecase interface {
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}
