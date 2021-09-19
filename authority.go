package goca

import "context"

type Authority interface {
	CanIssue(ctx context.Context, domain string) (bool, error)
	Order(ctx context.Context, domain string) 
}

type Order interface {
	
}