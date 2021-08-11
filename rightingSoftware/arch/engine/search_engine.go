package engine

import (
	"context"

	"github.com/emahiro/il/rightingSoftware/arch/resource"
)

type UserSearchEngine struct {
	ua resource.UserAccess
}

func NewUserSearchEngine(ua resource.UserAccess) UserSearchEngine {
	return UserSearchEngine{
		ua: ua,
	}
}

func (e *UserSearchEngine) SearchByMail(ctx context.Context, email string) (*resource.User, error) {
	return e.ua.GetByMail(ctx, email)
}
