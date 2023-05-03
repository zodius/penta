package handler

import (
	"net/http"
	"penta/model"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	DirListService model.DirListService
}

func AddHandler(r gin.IRouter, DirListService model.DirListService) {
	h := ServiceHandler{
		DirListService: DirListService,
	}

	r.POST("/api/upload", h.UploadUrlList)
}

func (h *ServiceHandler) UploadUrlList(c *gin.Context) {
	url_file, err := c.FormFile("url")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f, err := url_file.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.DirListService.ImportUrlList(f)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "OK")
}
