package graph

import (
	"testgqlgen/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Users struct {
	List   []*model.User
	LastID *int
}

type Resolver struct {
	UserList *Users
	AppList  []*model.App
}
