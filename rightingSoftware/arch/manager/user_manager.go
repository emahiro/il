package manager

import (
	"context"
	"errors"

	"github.com/emahiro/il/rightingSoftware/arch/engine"
)

type UserManager struct {
	ure engine.UserRegisterEngine
	use engine.UserSearchEngine
	me  engine.MailEngine
}

func NewUserManager(
	ure engine.UserRegisterEngine,
	use engine.UserSearchEngine,
	me engine.MailEngine,
) *UserManager {
	return &UserManager{
		ure: ure,
		use: use,
		me:  me,
	}
}

type UserRegisterParams struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Age  int64  `json:"age"`
}

func (m *UserManager) Register(ctx context.Context, p UserRegisterParams) error {
	user, err := m.use.SearchByMail(ctx, p.Mail)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("登録されたメールアドレスはすでにゆーざーが存在します。")
	}
	if err := m.ure.Register(ctx, p.Name, p.Age); err != nil {
		return err
	}
	return m.me.Send(ctx, p.Mail, struct{}{})
}

// 以下は Sample
func (m *UserManager) Update(ctx context.Context) error {
	return nil
}

func (m *UserManager) GetSelf(ctx context.Context) error {
	return nil
}
