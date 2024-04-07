package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type listGroupMembersRequest struct {
	GroupID int32 `uri:"groupId" binding:"required,min=1"`
}

func (s *Server) listGroupMembers(ctx *gin.Context) {
	var req listGroupMembersRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	members, err := s.store.ListMembersOfGroup(ctx, req.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, members)
}
