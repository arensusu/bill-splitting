package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

func (s *Server) GetDiscordGroup(ctx context.Context, req *proto.GetDiscordGroupRequest) (*proto.GetDiscordGroupResponse, error) {
	if req.DiscordChannel == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid discord channel")
	}

	discordGroup, err := s.store.GetGroupByDiscordChannel(req.DiscordChannel)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &proto.GetDiscordGroupResponse{
		Id:      uint32(discordGroup.ID),
		GroupId: uint32(discordGroup.ID),
		Name:    discordGroup.Name,
	}, nil
}

func (s *Server) SetDiscordGroup(ctx context.Context, req *proto.SetDiscordGroupRequest) (*proto.SetDiscordGroupResponse, error) {
	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or discord channel")
	}

	discordGroup, err := s.store.GetGroup(uint(req.GroupId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err = s.store.UpdateGroup(&model.Group{
		Model: gorm.Model{
			ID: discordGroup.ID,
		},
		DiscordChannel: req.DiscordChannel,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &proto.SetDiscordGroupResponse{
		Id:             uint32(discordGroup.ID),
		GroupId:        uint32(discordGroup.ID),
		DiscordChannel: discordGroup.DiscordChannel,
	}, nil
}

func (s *Server) CreateDiscordGroup(ctx context.Context, req *proto.CreateDiscordGroupRequest) (*proto.CreateDiscordGroupResponse, error) {
	if req.DiscordChannel == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid discord channel")
	}

	err := s.store.CreateGroup(&model.Group{
		Name:           req.Name,
		Currency:       req.Currency,
		DiscordChannel: req.DiscordChannel,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	group, err := s.store.GetGroupByDiscordChannel(req.DiscordChannel)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	user, err := s.store.GetUserByDiscordID(req.DiscordId)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	s.store.CreateMember(&model.Member{
		GroupID: group.ID,
		UserID:  user.ID,
	})

	return &proto.CreateDiscordGroupResponse{
		Id:             uint32(group.ID),
		Name:           req.Name,
		Currency:       req.Currency,
		DiscordChannel: req.DiscordChannel,
	}, nil
}

func (s *Server) ListGroupsOfUser(ctx context.Context, req *proto.ListGroupsOfUserRequest) (*proto.ListGroupsOfUserResponse, error) {
	user, err := s.store.GetUserByDiscordID(req.DiscordId)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	groups, err := s.store.ListGroupsByUserID(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	var protoGroups []*proto.Group
	for _, group := range groups {
		protoGroups = append(protoGroups, &proto.Group{
			Id:   uint32(group.ID),
			Name: group.Name,
		})
	}

	return &proto.ListGroupsOfUserResponse{
		Groups: protoGroups,
	}, nil
}
