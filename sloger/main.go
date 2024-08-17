package main

import (
	"context"

	"github.com/emahiro/il/sloger/sloger"
)

func main() {
	sloger.New()
	ctx := context.Background()
	sloger.Infof(ctx, "Hello, %s", "world")
}
