package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
)

type Resolver struct{
	UserList []*User
	MaxID *int
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	*r.MaxID = *r.MaxID + 1
	newUser := &User {
		ID: *r.MaxID,
		Name: input.Name,
		Age: input.Age,
	}
	r.UserList = append(r.UserList, newUser)
	return newUser, nil
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, id int) (*User, error) {
	for u, user := range r.UserList {
		if user.ID == id {
			r.UserList = append(r.UserList[:u], r.UserList[u+1:]...)
			return user, nil
		}
	}
	return nil, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	if len(r.UserList) == 0 {
		return nil, fmt.Errorf("No users stored on database")
	}
	return r.UserList, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
