package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker *token.JWTMaker
}

func NewServer(store db.Store, tokenMaker *token.JWTMaker) *Server {
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	// Externalize CORS configuration
	config := cors.Config{
		AllowOrigins: []string{"http://node-dev:3000", "http://localhost"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}
	router.Use(cors.New(config))

	// Setup API versioning
	api := router.Group("/api/v1")
	{
		api.GET("/auth/:provider", server.auth)
		api.GET("/auth/:provider/callback", server.authCallback)
		api.POST("/auth/linebot", server.authLineBot)

		api.Static("/images", "/var/images")

		// Authenticated routes
		authRoutes := api.Group("")
		authRoutes.Use(authMiddleware(server.tokenMaker))

		authRoutes.GET("/invites/:code", server.acceptGroupInvitation)

		authRoutes.DELETE("/settlements/:payerId/:payeeId", server.completeSettlement)

		groupRoutes := authRoutes.Group("/groups")

		groupRoutes.POST("", server.createGroup)
		groupRoutes.GET("/:groupId", server.getGroup)
		groupRoutes.GET("", server.listGroups)

		groupRoutes.POST("/:groupId/invites", server.createGroupInvitation)

		groupRoutes.GET("/:groupId/members", server.listGroupMembers)

		groupRoutes.POST("/:groupId/expenses", server.createExpense)
		groupRoutes.GET("/:groupId/expenses", server.listExpenses)
		//groupRoutes.GET("/:groupId/expenses/summary", server.listExpensesSummary)

		groupRoutes.PUT("/:groupId/settlements", server.replaceSettlement)
	}

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
