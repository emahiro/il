package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"os"
	"strings"
	"sync"
)

func main() {
	logger := slog.New(NewLtsvHandler(os.Stderr, nil))
	logger.Info("foo:aaa\tbar:bbb")
}

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return bytes.NewBuffer(b)
	},
}

func newBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

type LtsvHandler struct {
	ltsv bool
	mu   sync.Mutex
	w    io.Writer
	opts *slog.HandlerOptions
}

func NewLtsvHandler(w io.Writer, opts *slog.HandlerOptions) *LtsvHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &LtsvHandler{
		ltsv: true,
		w:    w,
		opts: opts,
	}
}

func (h *LtsvHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	return level >= minLevel
}

func (h *LtsvHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := newBuffer()
	msg := r.Message
	for k, v := range ltsvSeparator(msg) {
		buf.WriteString(fmt.Sprintf("%s=%s ", k, v))
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.w.Write(buf.Bytes())
	return nil
}

func ltsvSeparator(msg string) iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, v := range strings.Split(msg, "\t") {
			parts := strings.Split(v, ":")
			if len(parts) != 2 {
				continue
			}
			if !yield(parts[0], parts[1]) {
				return
			}
		}
	}
}

// TODO: Implement the rest of the slog.Handler interface
func (h *LtsvHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return slog.Default().Handler()
}

// TODO Implement the rest of the slog.Handler interface
func (h *LtsvHandler) WithGroup(name string) slog.Handler {
	return slog.Default().Handler()
}
