package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	httpS "gateway/internal/delivery/http"
	"gateway/pb"

	"github.com/danh996/go-shop-kit/token"

	"github.com/danh996/go-shop-kit/grpc_client"
	"github.com/danh996/go-shop-kit/grpc_server"

	"github.com/urfave/cli"
)

type srv struct {
	cfg *Config

	authClient     pb.AuthServiceClient
	userClient     pb.UserServiceClient
	productClient  pb.ProductServiceClient
	categoryClient pb.CategoryServiceClient
	searchClient   pb.SearchServiceClient

	authenticator token.Authenticator
	grpcServer    *grpc_server.GRPCServer
}

func (s *srv) loadGRPCClient() error {
	var err error

	authConn, err := grpc_client.NewGRPCClient(s.cfg.AuthEndpoint).Dial()
	if err != nil {
		return err
	}
	s.authClient = pb.NewAuthServiceClient(authConn)

	userConn, err := grpc_client.NewGRPCClient(s.cfg.AuthEndpoint).Dial()
	if err != nil {
		return err
	}
	s.userClient = pb.NewUserServiceClient(userConn)

	productConn, err := grpc_client.NewGRPCClient(s.cfg.ProductEndpoint).Dial()
	if err != nil {
		return err
	}
	s.productClient = pb.NewProductServiceClient(productConn)

	categoryConn, err := grpc_client.NewGRPCClient(s.cfg.CategoryEndpoint).Dial()
	if err != nil {
		return err
	}

	s.categoryClient = pb.NewCategoryServiceClient(categoryConn)

	searchConn, err := grpc_client.NewGRPCClient(s.cfg.SearchEndpoint).Dial()
	if err != nil {
		return err
	}
	s.searchClient = pb.NewSearchServiceClient(searchConn)

	log.Println("load GRPC CLIENT success")
	return nil
}

func (s *srv) loadGRPCServer() error {
	s.grpcServer = grpc_server.NewGRPCServer(s.cfg.ServiceName, s.cfg.HTTP.Host, s.cfg.HTTP.Port)

	s.grpcServer.InitServer()

	return nil
}

func (s *srv) startGRPCServer() error {
	ln, err := net.Listen("tcp", s.cfg.GRPC.String())
	if err != nil {
		return err
	}

	log.Printf("GRPC server listening in port: %v", s.cfg.GRPC.Port)

	if err := s.grpcServer.Server.Serve(ln); err != nil {
		return err
	}
	return nil
}

func (s *srv) startHTTPServer() error {
	ctx := context.Background()

	mux, err := httpS.NewHTTPHandler(
		ctx,
		s.authClient,
		s.productClient,
		s.categoryClient,
		s.userClient,
		s.authenticator,
		s.searchClient,
	)
	if err != nil {
		return err
	}
	log.Println("start http success in port", s.cfg.HTTP.Port)
	return http.ListenAndServe(":"+s.cfg.HTTP.Port, mux)
}

// func (s *srv) loadLogger() error {
// 	return nil
// }

func (s *srv) loadConfig(ctx *cli.Context) error {
	s.cfg = &Config{
		HTTP: ConnAddress{
			Host: ctx.GlobalString(HTTPHostFlag.GetName()),
			Port: ctx.GlobalString(HTTPPortFlag.GetName()),
		},
		GRPC: ConnAddress{
			Host: ctx.GlobalString(GRPCHostFlag.GetName()),
			Port: ctx.GlobalString(GRPCPortFlag.GetName()),
		},

		ServiceName:     ctx.GlobalString(ServiceNameFlag.GetName()),
		AuthEndpoint:    ctx.GlobalString(AuthEndpointFlag.GetName()),
		ProductEndpoint: ctx.GlobalString(ProductEndpointFlag.GetName()),
		SearchEndpoint:  ctx.GlobalString(SearchEndpointFlag.GetName()),
		TokenKey:        ctx.GlobalString(TokenKeyFlag.GetName()),
	}
	return nil
}

func (s *srv) loadAuthenticator() error {

	fmt.Println(s.cfg.TokenKey)

	var err error
	s.authenticator, err = token.NewJWTAuthenticator(s.cfg.TokenKey, 15*24*time.Hour)
	return err
}
