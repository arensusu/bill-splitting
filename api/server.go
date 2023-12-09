package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"

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

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.POST("/users/login", server.loginUser)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("/groups", server.createGroup)
	authRouter.GET("/groups/:id", server.getGroup)

	authRouter.POST("/group-members", server.createGroupMember)
	authRouter.GET("/group-members/:groupId", server.listGroupMembers)

	authRouter.POST("/expenses", server.createExpense)
	authRouter.GET("/expenses/:groupId", server.listExpenses)

	authRouter.PUT("/settlements", server.replaceSettlement)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
