package util

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"

	"github.com/aplulu/iapproxy/internal/config"
)

func GetTokenSource(ctx context.Context, audience string) (oauth2.TokenSource, error) {
	var tokenSource oauth2.TokenSource

	if config.GoogleImpersonateServiceAccount() != "" {
		ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
			TargetPrincipal: config.GoogleImpersonateServiceAccount(),
			Scopes:          []string{},
			Delegates:       []string{},
		}, option.WithAudiences(audience))
		if err != nil {
			return nil, fmt.Errorf("util.GetTokenSource: failed to create impersonated token source: %w", err)
		}

		tokenSource = ts
	} else {
		ts, err := idtoken.NewTokenSource(ctx, audience)
		if err != nil {
			return nil, fmt.Errorf("util.GetTokenSource: failed to create token source: %w", err)
		}

		tokenSource = ts
	}

	return tokenSource, nil
}
