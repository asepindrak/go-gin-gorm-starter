package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{ svc Service }

func NewHandler(s Service) *Handler { return &Handler{svc: s} }

type createReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type updateReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) Register(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/users", h.List)
		api.POST("/users", h.Create)
		api.GET("/users/:id", h.Get)
		api.PUT("/users/:id", h.Update)
		api.DELETE("/users/:id", h.Delete)
	}
}

func (h *Handler) Create(c *gin.Context) {
}
