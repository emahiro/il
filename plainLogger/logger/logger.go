package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var now = func() time.Time { return time.Now() }

type logger struct {
	mu  sync.Mutex
	out io.Writer
	buf []byte
}

func new(out io.Writer) *logger {
	return &logger{out: out}
}

func (l *logger) output(s string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := now()
	y, m, d := now.Date()
	hour, min, sec := now.Clock()

	line := fmt.Sprintf("%d/%d/%d %d:%d:%d %s", y, m, d, hour, min, sec, s)
	l.buf = l.buf[:0]
	l.buf = append(l.buf, line...)
	if len(l.buf) == 0 || line[len(line)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	if _, err := l.out.Write(l.buf); err != nil {
		return err
	}
	return nil
}

var std = new(os.Stdout)

func SetOutput(out io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()

	std.out = out
}

func Debugf(format string, a ...interface{}) {
	std.output(fmt.Sprintf(format, a...))
}
