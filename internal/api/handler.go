package api

import (
	"context"
	"devops-portfolio/app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) RootCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok v6"})
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok v6"})
}

func (h *Handler) CreateUser(c *gin.Context) {
	// 🔥 age 추가
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreateUser(context.Background(), req.Name, req.Email, req.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	u, err := h.userService.GetUser(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
