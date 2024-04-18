package gapi

import (
	"fmt"
	
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
	"github.com/techschool/simplebank/pb"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimplebankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
