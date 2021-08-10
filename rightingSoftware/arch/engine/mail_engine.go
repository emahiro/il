package engine

import (
	"context"
	"encoding/json"

	"github.com/emahiro/il/rightingSoftware/arch/util"
)

type MailEngine struct{}

func NewMailEngine() MailEngine {
	return MailEngine{}
}

func (m *MailEngine) Send(ctx context.Context, address string, data struct{}) error {
	body := struct {
		Address string   `json:"address"`
		Data    struct{} `json:"data"`
	}{
		Address: address,
		Data:    struct{}{},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return util.EnqueueMail(ctx, b)
}
