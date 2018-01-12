package context

import (
	"context"

	"../models"
)

type privateKey string

const (
	userKey privateKey = "user"
)

// WithUser returns a new Context with key "user"
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User Returns a User type given a Context
func User(ctx context.Context) *models.User {
	if temp := ctx.Value(userKey); temp != nil {
		if user, ok := temp.(*models.User); ok {
			return user
		}
	}
	return nil
}
