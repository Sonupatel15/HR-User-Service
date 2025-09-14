package server

import (
	"HR-User-Service/internal/models"
	"HR-User-Service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service *services.RoleService
}

func NewRoleHandler(service *services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req models.CreateRoleRequest

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.CreateRole(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, role)
}
