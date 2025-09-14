package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"HR-User-Service/internal/repository"
	"HR-User-Service/internal/services"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	// role endpoints
	roleRepo := repository.NewRoleRepository(s.db) // ✅ pass db
	roleService := services.NewRoleService(roleRepo)
	roleHandler := NewRoleHandler(roleService)

	r.POST("/roles", roleHandler.CreateRole)
	r.GET("/roles/:id", roleHandler.GetRoleByID)
	r.GET("/roles/name/:name", roleHandler.GetRoleByName)
	r.GET("/roles/exists", roleHandler.RoleExists)

	// user endpoints - ADD THESE
	userRepo := repository.NewUserRepository(s.db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.GET("/users/email/:email", userHandler.GetUserByEmail)
	r.GET("/users/exists", userHandler.UserExists)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
