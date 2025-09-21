package auth

import (
	"context"
	"fmt"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

func NewZitadelAuthorizer(ctx context.Context, domain, keypath string) (*authorization.Authorizer[*oauth.IntrospectionContext], error) {
	var zitadelOptions []zitadel.Option
	// if insecure {
	// 	zitadelOptions = append(zitadelOptions, zitadel.WithInsecure(domain))
	// }

	instance := zitadel.New(domain, zitadelOptions...)

	//authorizer, err := authorization.New(ctx, instance, oauth.WithIntrospection[*oauth.IntrospectionContext](oauth.ClientIDSecretIntrospectionAuthentication(clientID, clientSecret)))
	authorizer, err := authorization.New(ctx, instance, oauth.DefaultAuthorization(keypath))
	if err != nil {
		return nil, fmt.Errorf("failed to create authorizer: %w", err)
	}

	return authorizer, nil
}
