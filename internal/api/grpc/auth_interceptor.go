package grpcapi

import (
	"context"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

type AuthInterceptor struct {
	authorizer authorization.Authorizer[*oauth.IntrospectionContext]
}

func NewAuthInterceptor(authorizer authorization.Authorizer[*oauth.IntrospectionContext]) *AuthInterceptor {
	return &AuthInterceptor{authorizer: authorizer}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		userID, err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "userID", userID)
		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	claims, err := i.authorizer.CheckAuthorization(ctx, values[0])
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return claims.UserID(), nil
}
