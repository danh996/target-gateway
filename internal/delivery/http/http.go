package http

import (
	"context"
	"net/http"

	"gateway/pb"

	"github.com/danh996/go-shop-kit/token"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// NewHTTPHandler ...
func NewHTTPHandler(
	ctx context.Context,
	authClient pb.AuthServiceClient,
	productClient pb.ProductServiceClient,
	categoryClient pb.CategoryServiceClient,
	userClient pb.UserServiceClient,
	authenticator token.Authenticator,
	searchClient pb.SearchServiceClient,

) (http.Handler, error) {
	mux := runtime.NewServeMux()

	if err := pb.RegisterAuthServiceHandlerClient(ctx, mux, authClient); err != nil {
		return nil, err
	}

	if err := pb.RegisterUserServiceHandlerClient(ctx, mux, userClient); err != nil {
		return nil, err
	}

	if err := pb.RegisterProductServiceHandlerClient(ctx, mux, productClient); err != nil {
		return nil, err
	}

	if err := pb.RegisterCategoryServiceHandlerClient(ctx, mux, categoryClient); err != nil {
		return nil, err
	}

	if err := pb.RegisterSearchServiceHandlerClient(ctx, mux, searchClient); err != nil {
		return nil, err
	}

	return Authorized(authenticator, userClient, mux), nil

	//return cors.AllowAll().Handler(Authorized(authenticator, userClient, mux)), nil

	//return mux, nil
}
