package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gateway/pb"
	"github.com/danh996/go-shop-kit/token"

	"github.com/danh996/go-shop-kit/requestinfo"

	"google.golang.org/grpc/metadata"
)

var whiteList = []string{
	"/auth",
	"/provinces",
	"/districts",
	"/wards",
}

func Authorized(
	authenticator token.Authenticator,
	userClient pb.UserServiceClient,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range whiteList {
			if strings.Contains(r.URL.Path, v) {
				next.ServeHTTP(w, r)
				return
			}
		}

		ctx := r.Context()
		authorization := r.Header.Get(requestinfo.Authorization)

		if authorization == "" {
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"message": "missing authorization",
			})
			return
		}
		bearerToken := strings.Split(authorization, requestinfo.Bearer+" ")
		if len(bearerToken) < 2 {
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"message": "authorization is invalid",
			})
			return
		}
		token := bearerToken[1]
		payload, err := authenticator.Verify(token)
		if err != nil {
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"message": fmt.Errorf("unable to verify token: %w", err).Error(),
			})
			return
		}
		user, err := userClient.GetUserByID(ctx, &pb.GetUserByIDRequest{
			UserID: payload.UserID,
		})

		if err != nil {
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"message": fmt.Errorf("unable to get user by id: %w", err).Error(),
			})
			return
		}
		// authorized here

		ctx = context.WithValue(ctx, requestinfo.Info{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func AppendRequestMetadata(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}
	// info, ok := ctx.Value(requestinfo.Info{}).(*pb.GetUserByIDResponse)

	// if !ok {
	// 	return md
	// }

	// b, err := json.Marshal(&grpc_mapping.Info{
	// 	ID:            info.UserID,
	// 	Fullname:      info.Fullname,
	// 	Phone:         info.Phone,
	// 	Email:         info.Email,
	// 	Job:           info.Job,
	// 	Age:           info.Age,
	// 	Role:          int(info.Role),
	// 	Sex:           int(info.Sex),
	// 	ProvinceID:    info.ProvinceID,
	// 	DistrictID:    info.DistrictID,
	// 	WardID:        info.WardID,
	// 	LocationScore: info.LocationScore,
	// 	IdentityCard:  info.IdentityCard,
	// })
	// if err == nil {
	// 	md.Append(grpc_mapping.InfoKey, string(b))
	// }

	return md
}
