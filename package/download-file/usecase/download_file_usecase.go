package usecase

import (
	"github.com/novalwardhana/golang-boilerplate/package/download-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/download-file/repository"
)

type usecase struct {
	repository repository.Repository
}

type Usecase interface {
	GetFileInfo(filename string) <-chan model.Result
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetFileInfo(filename string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		/* Get file info */
		fileInfo := <-u.repository.GetFileInfo(filename)
		if fileInfo.Error != nil {
			output <- model.Result{Error: fileInfo.Error}
			return
		}

		output <- model.Result{Data: fileInfo.Data}
	}()
	return output
}
