package handler

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

const ExtCSV string = "csv"
const EXTPDF string = "pdf"
const ExtExcel string = "xlsx"
const ExtZIP string = "zip"

func (h *Handler) Mount(g *echo.Group) {
	g.POST("/csv", h.uploadCSV)
	g.POST("/pdf", h.uploadPDF)
	g.POST("/excel", h.uploadExcel)
	g.POST("/zip", h.uploadZIP)
	g.POST("/csv-to-database", h.uploadCSVToDatabase)
	g.POST("/multiple", h.uploadMultipleFile)
}

func (h *Handler) uploadCSV(c echo.Context) error {
	var file *multipart.FileHeader
	var err error
	var response model.Response

	/* Get file form */
	if file, err = c.FormFile("file"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file parameter"
		return c.JSON(http.StatusOK, response)
	}

	/* File validation */
	filenameArr := strings.Split(file.Filename, ".")
	if len(filenameArr) < 2 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File not valid"
		return c.JSON(http.StatusOK, response)
	}
	if filenameArr[1] != ExtCSV {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File must CSV format"
		return c.JSON(http.StatusOK, response)
	}
	fileExt := filenameArr[1]

	/* Process upload file */
	uploadResult := <-h.usecase.UploadFile(file, fileExt)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "File csv successfully uploaded"
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) uploadPDF(c echo.Context) error {
	var file *multipart.FileHeader
	var response model.Response
	var err error

	/* Get file form */
	if file, err = c.FormFile("file"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file parameter"
		return c.JSON(http.StatusOK, response)
	}

	/* File validation */
	filenameArr := strings.Split(file.Filename, ".")
	if len(filenameArr) < 2 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File not valid"
		return c.JSON(http.StatusOK, response)
	}
	if filenameArr[1] != EXTPDF {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File must PDF format"
		return c.JSON(http.StatusOK, response)
	}
	fileExt := filenameArr[1]

	/* Process upload file */
	uploadResult := <-h.usecase.UploadFile(file, fileExt)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "File pdf successfully uploaded"
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) uploadExcel(c echo.Context) error {
	var file *multipart.FileHeader
	var response model.Response
	var err error

	/* Get file form */
	if file, err = c.FormFile("file"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file parameter"
		return c.JSON(http.StatusOK, response)
	}

	/* File validation */
	filenameArr := strings.Split(file.Filename, ".")
	if len(filenameArr) < 2 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File not valid"
		return c.JSON(http.StatusOK, response)
	}
	if filenameArr[1] != ExtExcel {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File must xlsx format"
		return c.JSON(http.StatusOK, response)
	}
	fileExt := filenameArr[1]

	/* Process upload file */
	uploadResult := <-h.usecase.UploadFile(file, fileExt)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "File excel successfully uploaded"
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) uploadCSVToDatabase(c echo.Context) error {
	var file *multipart.FileHeader
	var err error
	var response model.Response

	/* Get file form */
	if file, err = c.FormFile("file"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file parameter"
		return c.JSON(http.StatusOK, response)
	}

	/* File validation */
	filenameArr := strings.Split(file.Filename, ".")
	if len(filenameArr) < 2 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File not valid"
		return c.JSON(http.StatusOK, response)
	}
	if filenameArr[1] != ExtCSV {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File must CSV format"
		return c.JSON(http.StatusOK, response)
	}

	/* Upload csv to database */
	uploadResult := <-h.usecase.UploadCSVToDatabase(file)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "Successfully export to database"
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) uploadZIP(c echo.Context) error {
	var file *multipart.FileHeader
	var err error
	var response model.Response

	/* Get file form */
	if file, err = c.FormFile("file"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file parameter"
		return c.JSON(http.StatusOK, response)
	}

	/* File validation */
	filenameArr := strings.Split(file.Filename, ".")
	if len(filenameArr) < 2 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File not valid"
		return c.JSON(http.StatusOK, response)
	}
	if filenameArr[1] != ExtZIP {
		response.StatusCode = http.StatusBadRequest
		response.Message = "File must CSV format"
		return c.JSON(http.StatusOK, response)
	}

	/* Process upload file */
	uploadResult := <-h.usecase.UploadFileZIP(file)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "File zip successfully uploaded"
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) uploadMultipleFile(c echo.Context) error {
	var uploadFiles []model.UploadFile
	var response model.Response

	/* Get file1 form */
	if file, err := c.FormFile("file1"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file1 parameter: " + err.Error()
		return c.JSON(http.StatusOK, response)
	} else {
		filenameArr := strings.Split(file.Filename, ".")
		if len(filenameArr) < 2 {
			response.StatusCode = http.StatusBadRequest
			response.Message = "File1 format is not valid"
			return c.JSON(http.StatusOK, response)
		}
		uploadFile := model.UploadFile{
			File:    file,
			FileExt: filenameArr[1],
		}
		uploadFiles = append(uploadFiles, uploadFile)
	}

	/* Get file2 form */
	if file, err := c.FormFile("file2"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file2 parameter: " + err.Error()
		return c.JSON(http.StatusOK, response)
	} else {
		filenameArr := strings.Split(file.Filename, ".")
		if len(filenameArr) < 2 {
			response.StatusCode = http.StatusBadRequest
			response.Message = "File2 format is not valid"
			return c.JSON(http.StatusOK, response)
		}
		uploadFile := model.UploadFile{
			File:    file,
			FileExt: filenameArr[1],
		}
		uploadFiles = append(uploadFiles, uploadFile)
	}

	/* Get file3 form */
	if file, err := c.FormFile("file3"); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid form file2 parameter: " + err.Error()
		return c.JSON(http.StatusOK, response)
	} else {
		filenameArr := strings.Split(file.Filename, ".")
		if len(filenameArr) < 2 {
			response.StatusCode = http.StatusBadRequest
			response.Message = "File3 format is not valid"
		}
		uploadFile := model.UploadFile{
			File:    file,
			FileExt: filenameArr[1],
		}
		uploadFiles = append(uploadFiles, uploadFile)
	}

	/* Upload process */
	uploadResult := <-h.usecase.UploadMultipleFile(uploadFiles)
	if uploadResult.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = uploadResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = http.StatusOK
	response.Message = "Files successfully uploaded"
	return c.JSON(http.StatusOK, response)
}
