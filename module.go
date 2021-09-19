package goca

import (
	"context"
	"fmt"

	"gopkg.zouai.io/colossus/clog"
	"gopkg.zouai.io/goca/authority_acme"
	"gopkg.zouai.io/goca/config"
)

type GoCA struct {
	logger clog.LoggerInterface
	config *config.Config

	authorities map[string]Authority
}

func New(ctx context.Context, config *config.Config) (*GoCA, error) {
	m := &GoCA{
		logger:      clog.SubLoggerWithPrefix(ctx, "GoCA"),
		config:      config,
		authorities: map[string]Authority{},
	}
	for authorityName, authority := range config.Authorities {
		if authority.RootFile != nil {
			// Do something for RootCA
		}
		if authority.ACME != nil {
			acmeClient, err := authority_acme.New(ctx, authorityName, authority.ACME)
			if err != nil {
				return nil, fmt.Errorf("error creating acme client for authority '%s': %w", authorityName, err)
			}
			m.authorities[authorityName] = acmeClient
		}
	}
	return m, nil
}
