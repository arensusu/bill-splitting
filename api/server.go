package api

import (
	db "bill-splitting/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)

	router.POST("/groups", server.createGroup)
	router.GET("/groups/:id", server.getGroup)

	router.POST("/group-members", server.createGroupMember)
	router.GET("/group-members/:groupId", server.listGroupMembers)

	router.POST("/expenses", server.createExpense)

	router.POST("/settlements", server.createSettlement)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
