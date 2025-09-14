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

// POST /roles
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req models.CreateRoleRequest
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

// GET /roles/:id
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.service.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// GET /roles/name/:name
func (h *RoleHandler) GetRoleByName(c *gin.Context) {
	name := c.Param("name")
	role, err := h.service.GetRoleByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// GET /roles/exists?id=xxx OR /roles/exists?name=yyy
func (h *RoleHandler) RoleExists(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	var exists bool
	var err error

	if id != "" {
		exists, err = h.service.ExistsByID(id)
	} else if name != "" {
		exists, err = h.service.ExistsByName(name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id or name query param required"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
