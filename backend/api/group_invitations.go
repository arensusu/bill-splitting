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
	GroupID int32 `uri:"groupId" binding:"required,min=1"`
}

func (s *Server) createGroupInvitation(ctx *gin.Context) {
	var req createGroupInvitationParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := ctx.MustGet("payload").(*token.JWTPayload)

	member, err := s.store.GetMembership(ctx, db.GetMembershipParams{
		GroupID: req.GroupID,
		UserID:  payload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if payload.UserID != member.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
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

type acceptGroupInvitationRequest struct {
	Code string `uri:"code" binding:"required"`
}

func (s *Server) acceptGroupInvitation(ctx *gin.Context) {
	var req acceptGroupInvitationRequest
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
