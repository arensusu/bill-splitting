package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLineGroup(ctx context.Context, req *proto.CreateLineGroupRequest) (*proto.CreateLineGroupResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	if req.LineId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid line id")
	}

	group, err := s.store.CreateGroupTx(ctx, db.CreateGroupTxParams{
		Name:   req.Name,
		UserID: payload.UserID,
		LineId: req.LineId,
	})
	if err != nil {
		return nil, err
	}

	return &proto.CreateLineGroupResponse{
		Id:     group.ID,
		LineId: group.LineID.String,
		Name:   group.Name,
	}, nil
}

func (s *Server) AddGroupMember(ctx context.Context, req *proto.AddGroupMemberRequest) (*proto.AddGroupMemberResponse, error) {
	_, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	if req.GroupId == 0 || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or user id")
	}

	member, err := s.store.CreateMember(ctx, db.CreateMemberParams{
		GroupID: req.GroupId,
		UserID:  req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &proto.AddGroupMemberResponse{
		Id:      member.ID,
		GroupId: member.GroupID,
		UserId:  member.UserID,
	}, nil
}
