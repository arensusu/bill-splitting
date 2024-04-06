package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createGroupMemberRequest struct {
	Code string `json:"code" binding:"required"`
}

func (s *Server) createGroupMember(ctx *gin.Context) {
	var req createGroupMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invite, err := s.store.GetGroupInvitation(ctx, req.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid code")})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = s.store.DeleteGroupInvitation(ctx, req.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payload := ctx.MustGet("payload").(*token.JWTPayload)

	member, err := s.store.CreateMember(ctx, db.CreateMemberParams{
		GroupID: invite.GroupID,
		UserID:  payload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, member)
}

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
