package api

import (
	db "bill-splitting/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createGroupMemberRequest struct {
	GroupID int64  `json:"groupId" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
}

func (s *Server) createGroupMember(ctx *gin.Context) {
	var req createGroupMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := s.store.CreateGroupMember(ctx, db.CreateGroupMemberParams{
		GroupID: req.GroupID,
		UserID:  req.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, member)
}

type listGroupMembersRequest struct {
	GroupID int64 `uri:"groupId" binding:"required,min=1"`
}

func (s *Server) listGroupMembers(ctx *gin.Context) {
	var req listGroupMembersRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	members, err := s.store.ListGroupMembers(ctx, req.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, members)
}
