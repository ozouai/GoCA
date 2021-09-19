package authority_acme

import (
	"context"
	"crypto"
	"fmt"

	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"gopkg.zouai.io/colossus/clog"
	"gopkg.zouai.io/goca/config"
)

type ACME struct {
	logger clog.LoggerInterface
	client *lego.Client
}
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}
func New(ctx context.Context, authorityName string, config *config.ACME) (*ACME, error) {
	m := &ACME{
		logger: clog.SubLoggerWithPrefix(ctx, "ACME").SubLoggerWithFields(ctx, map[string]interface{}{"authority": authorityName}),
	}
	myUser := MyUser{
		Email: "you@yours.com",
	}
	legoConfig := lego.NewConfig(&myUser)
	client, err := lego.NewClient(legoConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating acme client: %w", err)
	}
	m.client = client
	return m, nil
}

func (m *ACME) CanIssue(ctx context.Context, domain string) (bool, error) {
	
}