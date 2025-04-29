package api

import (
	"bill-splitting/model"
	"bill-splitting/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	store      model.Store
	router     *gin.Engine
	tokenMaker *token.JWTMaker
}

func NewServer(store model.Store, tokenMaker *token.JWTMaker) *Server {
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	// Externalize CORS configuration
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}
	router.Use(cors.New(config))

	// Setup API versioning
	api := router.Group("/api/v1")

	external := api.Group("/external")
	{
		external.GET("/discord/:discordChannel/expenses", server.listExpenses)
	}

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
