package usecase

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"time"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"
	"github.com/novalwardhana/golang-boilerplate/package/download-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/download-file/repository"
)

type usecase struct {
	repository repository.Repository
}

type Usecase interface {
	DownloadFile(filename string) <-chan model.Result
	DownloadFileZip(filename string) <-chan model.Result
	DownloadMultipleFile(filenames []string) <-chan model.Result
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

const DefaultFileLocation string = "/home/novalwardhana/golang-boilerplate/upload-file"

func (u *usecase) DownloadFile(filename string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		/* Get file info */
		fileInfo := <-u.repository.GetFileInfo(filename)
		if fileInfo.Error != nil {
			output <- model.Result{Error: fileInfo.Error}
			return
		}

		/* Check file is exist or not */
		file := fileInfo.Data.(model.File)
		if _, err := os.Stat(file.Directory + file.Name); os.IsNotExist(err) {
			output <- model.Result{Error: err}
			return
		}

		output <- model.Result{Data: fileInfo.Data}
	}()
	return output
}

func (u *usecase) DownloadFileZip(filename string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		/* Get file info */
		fileInfo := <-u.repository.GetFileInfo(filename)
		if fileInfo.Error != nil {
			output <- model.Result{Error: fileInfo.Error}
			return
		}

		/* Create download zip directory */
		fileLocation := os.Getenv(globalENV.GeneralFileDir)
		if len(fileLocation) <= 0 {
			fileLocation = DefaultFileLocation
		}
		fileDirZip := fileLocation + "/download-zip/"
		if _, err := os.Stat(fileDirZip); os.IsNotExist(err) {
			err := os.MkdirAll(fileDirZip, os.ModePerm)
			if err != nil {
				output <- model.Result{Error: err}
				return
			}
		}

		/* Create zip file */
		filenameZip := "ZIP_" + time.Now().Format("20060102150405") + "_" + "download_file.zip"
		fileZip, err := os.OpenFile(fileDirZip+filenameZip, os.O_WRONLY|os.O_CREATE, os.ModePerm) // Read file or create if not exist
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		defer fileZip.Close()
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		fileZipWrite := zip.NewWriter(fileZip)
		defer fileZipWrite.Close()

		/* Check file is exist or not */
		file := fileInfo.Data.(model.File)
		if _, err := os.Stat(file.Directory + file.Name); os.IsNotExist(err) {
			output <- model.Result{Error: err}
			return
		}

		/* Open file source */
		fileSrc, err := os.Open(file.Directory + file.Name) // Readonly
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		defer fileSrc.Close()

		/* Get file source info */
		fileSrcInfo, err := fileSrc.Stat()
		if err != nil {
			output <- model.Result{Error: err}
			return
		}

		/* Get file zip header */
		fileZipHeader, err := zip.FileInfoHeader(fileSrcInfo)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		// fileZipHeader.Method = zip.Deflate
		// Unactivated because some files format error when file has compressed

		/* Add file to zip */
		writer, err := fileZipWrite.CreateHeader(fileZipHeader)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		if _, err := io.Copy(writer, fileSrc); err != nil {
			output <- model.Result{Error: err}
			return
		}

		zip := model.Zip{
			Directory: fileDirZip,
			Name:      filenameZip,
		}

		output <- model.Result{Data: zip}
	}()
	return output
}

func (u *usecase) DownloadMultipleFile(filenames []string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		var totalError int

		/* Create download zip directory */
		fileLocation := os.Getenv(globalENV.GeneralFileDir)
		if len(fileLocation) <= 0 {
			fileLocation = DefaultFileLocation
		}
		fileDirZip := fileLocation + "/dowload-zip/"
		if _, err := os.Stat(fileDirZip); os.IsNotExist(err) {
			err := os.MkdirAll(fileDirZip, os.ModePerm)
			if err != nil {
				output <- model.Result{Error: err}
				return
			}
		}

		/* Create zip file */
		filenameZip := "ZIP_" + time.Now().Format("20060102150405") + "_" + "download_file.zip"
		fileZip, err := os.OpenFile(fileDirZip+filenameZip, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		defer fileZip.Close()
		fileZipWrite := zip.NewWriter(fileZip)
		defer fileZipWrite.Close()

		/* files iteration */
		for _, filename := range filenames {

			/* Get file info */
			fileInfo := <-u.repository.GetFileInfo(filename)
			if fileInfo.Error != nil {
				totalError++
				continue
			}
			file := fileInfo.Data.(model.File)

			/* Check file is exist or not */
			if _, err := os.Stat(file.Directory + file.Name); os.IsNotExist(err) {
				totalError++
				continue
			}

			/* Open file src */
			fileSrc, err := os.Open(file.Directory + file.Name)
			if err != nil {
				totalError++
				continue
			}

			/* Get file src information */
			fileSrcInfo, err := fileSrc.Stat()
			if err != nil {
				totalError++
				continue
			}

			/* Set file zip header */
			fileZipHeader, err := zip.FileInfoHeader(fileSrcInfo)
			if err != nil {
				totalError++
				continue
			}

			/* Append file to zip */
			writer, err := fileZipWrite.CreateHeader(fileZipHeader)
			if err != nil {
				totalError++
				continue
			}
			if _, err := io.Copy(writer, fileSrc); err != nil {
				totalError++
				continue
			}

			fileSrc.Close()
		}

		if totalError == len(filenames) {
			output <- model.Result{Error: errors.New("all file cannot zipped")}
			return
		}

		zip := model.Zip{
			Directory: fileDirZip,
			Name:      filenameZip,
		}
		output <- model.Result{Data: zip}

	}()
	return output
}
