package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/helper"
	"bill-splitting/token"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createGroupInvitationParams struct {
	GroupID int64 `json:"group_id"`
}

func (s *Server) createGroupInvitation(ctx *gin.Context) {
	var req createGroupInvitationParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := ctx.MustGet("payload").(*token.JWTPayload)

	_, err := s.store.GetGroupMember(ctx, db.GetGroupMemberParams{
		GroupID: req.GroupID,
		UserID:  payload.UserID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusForbidden, gin.H{"error": errors.New("user is not a member of the group")})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code := helper.RandomString(8)
	invite, err := s.store.CreateGroupInvitation(ctx, db.CreateGroupInvitationParams{
		Code:    code,
		GroupID: req.GroupID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, invite)
}

type getGroupInvitationRequest struct {
	Code string `uri:"code" binding:"required"`
}

func (s *Server) getGroupInvitation(ctx *gin.Context) {
	var req getGroupInvitationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
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

	ctx.JSON(http.StatusOK, invite)
}
