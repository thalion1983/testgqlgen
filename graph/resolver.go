package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
)

type Resolver struct{}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	panic("not implemented")
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
