package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker *token.JWTMaker
}

func NewServer(store db.Store, secretKey string) *Server {
	tokenMaker := token.NewJWTMaker(secretKey)
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://node-dev:3000"}
	router.Use(cors.New(config))

	router.GET("/auth/:provider", server.auth)
	router.GET("/auth/:provider/callback", server.authCallback)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("/groups", server.createGroup)
	authRouter.GET("/groups/:id", server.getGroup)
	authRouter.GET("/groups", server.listGroups)

	authRouter.POST("/groups/members", server.createGroupMember)
	authRouter.GET("/groups/members/:groupId", server.listGroupMembers)

	authRouter.POST("/expenses", server.createExpense)
	authRouter.GET("/expenses/:groupId", server.listExpenses)

	authRouter.PUT("/settlements", server.replaceSettlement)
	authRouter.DELETE("/settlements", server.completeSettlement)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
