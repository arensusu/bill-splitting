package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"bill-splitting/token"
)

type Server struct {
	proto.UnimplementedBillSplittingServer
	store      *model.Store
	tokenMaker *token.JWTMaker
}

func NewServer(store *model.Store, tokenMaker *token.JWTMaker) *Server {
	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
}
