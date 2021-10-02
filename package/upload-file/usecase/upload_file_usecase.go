package usecase

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/repository"
)

type usecase struct {
	repository repository.Repository
}

const DefaultFileLocation string = "/home/novalwardhana/golang-boilerplate/upload-file"

type Usecase interface {
	UploadFile(file *multipart.FileHeader, fileExt string) <-chan model.Result
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) UploadFile(file *multipart.FileHeader, fileExt string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		/* Choose file location */
		fileLocation := os.Getenv(globalENV.GeneralFileDir)
		if len(fileLocation) <= 0 {
			fileLocation = DefaultFileLocation
		}

		/* Create file directory */
		fileDir := fileLocation + "/" + fileExt + "/"
		if _, err := os.Stat(fileDir); os.IsNotExist(err) {
			err := os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				output <- model.Result{Error: err}
				return
			}
		}

		/* Compose file src */
		fileSrc, err := file.Open()
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		defer fileSrc.Close()

		/* Compose file target */
		filename := "upload" + fileExt + "_" + time.Now().Format("060102150405") + "_" + strings.ReplaceAll(file.Filename, " ", "_")
		fileTarget, err := os.OpenFile(fileDir+filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		defer fileTarget.Close()

		/* Save file to directory */
		if _, err := io.Copy(fileTarget, fileSrc); err != nil {
			output <- model.Result{Error: err}
			return
		}

		output <- model.Result{}
	}()
	return output
}
