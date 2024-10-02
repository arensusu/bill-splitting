package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLineGroup(ctx context.Context, req *proto.CreateLineGroupRequest) (*proto.CreateLineGroupResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.LineId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid line id")
	}

	group, err := s.store.CreateGroupTx(model.CreateGroupTxParams{
		Name:       req.Name,
		UserLineId: payload.UserID,
		LineId:     req.LineId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &proto.CreateLineGroupResponse{
		Id:     0,
		LineId: group.LineId,
		Name:   group.Name,
	}, nil
}

func (s *Server) GetLineGroup(ctx context.Context, req *proto.GetLineGroupRequest) (*proto.GetLineGroupResponse, error) {
	_, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.LineId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid line id")
	}

	group, err := s.store.GetGroupByLineID(req.LineId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &proto.GetLineGroupResponse{
		Id:     uint32(group.ID),
		LineId: group.LineId,
		Name:   group.Name,
	}, nil
}

func (s *Server) AddMembership(ctx context.Context, req *proto.AddMembershipRequest) (*proto.AddMembershipResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or user id")
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err = s.store.CreateMember(&model.Member{
		GroupID: uint(req.GroupId),
		UserID:  user.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &proto.AddMembershipResponse{
		Id:      0,
		GroupId: uint32(req.GroupId),
		UserId:  user.LineID,
	}, nil
}

func (s *Server) GetMembership(ctx context.Context, req *proto.GetMembershipRequest) (*proto.GetMembershipResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or user id")
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	member, err := s.store.GetMembership(uint(req.GroupId), user.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &proto.GetMembershipResponse{
		Id:      0,
		GroupId: uint32(member.GroupID),
		UserId:  user.LineID,
	}, nil
}
