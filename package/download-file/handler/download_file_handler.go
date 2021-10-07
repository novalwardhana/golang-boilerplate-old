package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/novalwardhana/golang-boilerplate/package/download-file/model"
	"github.com/novalwardhana/golang-boilerplate/package/download-file/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Mount(g *echo.Group) {
	g.GET("/:filename", h.downloadFile)
	g.GET("/zip/:filename", h.downloadFileZip)
	g.GET("/multiple", h.downloadMultipleFile)
}

func (h *Handler) downloadFile(c echo.Context) error {
	var response model.Response

	/* Get file info */
	filename := c.Param("filename")
	fileInfo := <-h.usecase.DownloadFile(filename)
	if fileInfo.Error != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = fileInfo.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	/* Download file */
	file := fileInfo.Data.(model.File)
	c.Attachment(file.Directory+file.Name, file.Name)
	return nil
}

func (h *Handler) downloadFileZip(c echo.Context) error {
	var response model.Response

	/* Get file info */
	filename := c.Param("filename")
	fileInfo := <-h.usecase.DownloadFileZip(filename)
	if fileInfo.Error != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = fileInfo.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	/* Download file */
	zip := fileInfo.Data.(model.Zip)
	return c.Attachment(zip.Directory+zip.Name, zip.Name)
}

func (h *Handler) downloadMultipleFile(c echo.Context) error {
	var payload model.MultipleFilePayload
	var response model.Response

	/* Payload validation */
	if err := c.Bind(&payload); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	if len(payload.Filenames) == 0 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Filenames must be filled in payload"
		return c.JSON(http.StatusOK, response)
	}

	/* Download process */
	downloadProcess := <-h.usecase.DownloadMultipleFile(payload.Filenames)
	if downloadProcess.Error != nil {
		response.StatusCode = http.StatusBadGateway
		response.Message = downloadProcess.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	zip := downloadProcess.Data.(model.Zip)
	return c.Attachment(zip.Directory+zip.Name, zip.Name)
}
