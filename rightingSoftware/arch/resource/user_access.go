package resource

import (
	"context"
	"database/sql"
)

type UserAccess interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByMail(ctx context.Context, mail string) (*User, error)
	Create(ctx context.Context, src *User) error
}

type userAccess struct{}

func NewUserAccess(db *sql.DB) UserAccess {
	return &userAccess{}
}

func (a *userAccess) GetByID(ctx context.Context, id int64) (*User, error) {
	return nil, nil
}

func (e *userAccess) GetByMail(ctx context.Context, mail string) (*User, error) {
	return nil, nil
}

func (a *userAccess) Create(ctx context.Context, src *User) error {
	return nil
}
