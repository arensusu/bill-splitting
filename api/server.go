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

	router.POST("/login", server.loginUser)

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)

	router.POST("/groups", server.createGroup)
	router.GET("/groups/:id", server.getGroup)

	router.POST("/group-members", server.createGroupMember)
	router.GET("/group-members/:groupId", server.listGroupMembers)

	router.POST("/expenses", server.createExpense)
	router.GET("/expenses/:groupId", server.listExpenses)

	router.PUT("/settlements", server.replaceSettlement)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
