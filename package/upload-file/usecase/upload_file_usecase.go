package usecase

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/repository"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repository repository.Repository
}

const DefaultFileLocation string = "/home/novalwardhana/golang-boilerplate/upload-file"

type Usecase interface {
	UploadFile(file *multipart.FileHeader, fileExt string) <-chan model.Result
	UploadCSVToDatabase(file *multipart.FileHeader) <-chan model.Result
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

		/* Save file information to database */
		filesize := strconv.Itoa(int(file.Size))
		fileInfo := model.File{
			Directory: fileDir,
			Name:      filename,
			Size:      filesize,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		saveFileInfoResult := <-u.repository.SaveFileInfo(fileInfo)
		if saveFileInfoResult.Error != nil {
			output <- model.Result{Error: saveFileInfoResult.Error}
			return
		}

		output <- model.Result{}
	}()
	return output
}

func (u *usecase) UploadCSVToDatabase(file *multipart.FileHeader) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		/* Choose file location */
		fileLocation := os.Getenv(globalENV.GeneralFileDir)
		if len(fileLocation) <= 0 {
			fileLocation = DefaultFileLocation
		}

		/* Create file directory */
		fileDir := fileLocation + "/csv-to-database/"
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

		/* Compose file target */
		filename := "uploadCSV_" + time.Now().Format("060102150405") + "_" + strings.ReplaceAll(file.Filename, " ", "_")
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

		/* Open file after upload */
		fileCSV, err := os.Open(fileDir + filename)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}

		/* Scan file after upload */
		var users []model.User
		scanner := bufio.NewScanner(fileCSV)
		for scanner.Scan() {
			line := scanner.Text()
			lineSplit := strings.Split(line, ",")

			encryptPassword, err := bcrypt.GenerateFromPassword([]byte(lineSplit[3]), model.PasswordHashCost)
			if err != nil {
				continue
			}
			user := model.User{
				Name:      lineSplit[0],
				Username:  lineSplit[1],
				Email:     lineSplit[2],
				Password:  string(encryptPassword),
				IsActive:  true,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}
			users = append(users, user)
		}
		if len(users) < 2 {
			output <- model.Result{Error: errors.New("data in csv not found")}
			return
		}
		CSVToDatabaseResult := <-u.repository.CSVToDatabase(users[1:])
		if CSVToDatabaseResult.Error != nil {
			output <- model.Result{Error: CSVToDatabaseResult.Error}
			return
		}

		output <- model.Result{}
	}()
	return output
}
