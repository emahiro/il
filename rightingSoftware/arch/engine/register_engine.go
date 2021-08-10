package engine

import (
	"context"

	"github.com/emahiro/il/rightingSoftware/arch/resource"
)

type UserRegisterEngine interface {
	Register(ctx context.Context, name string, age int64) error
}

type userRegisterEngine struct {
	ua resource.UserAccess
}

func NewRegisterEngine(ua resource.UserAccess) UserRegisterEngine {
	return &userRegisterEngine{
		ua: ua,
	}
}

func (e *userRegisterEngine) Register(ctx context.Context, name string, age int64) error {
	return e.ua.Create(ctx, &resource.User{Name: name, Age: age})
}
