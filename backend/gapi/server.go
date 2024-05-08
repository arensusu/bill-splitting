package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"bill-splitting/token"
)

type Server struct {
	proto.UnimplementedBillSplittingServer
	store      db.Store
	tokenMaker *token.JWTMaker
}

func NewServer(store db.Store, tokenMaker *token.JWTMaker) *Server {
	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
}
