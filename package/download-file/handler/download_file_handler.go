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
