package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Total int64 `json:"total,omitempty"`
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
}

type Envelope struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  *Meta       `json:"meta,omitempty"`
	Error string      `json:"error,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, data)
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Envelope{Data: data})
}

func Paginated(c *gin.Context, data interface{}, total int64, page, size int) {
	c.JSON(http.StatusOK, Envelope{Data: data, Meta: &Meta{Total: total, Page: page, Size: size}})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Envelope{Error: err.Error()})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Envelope{Error: msg})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Envelope{Error: err.Error()})
}
