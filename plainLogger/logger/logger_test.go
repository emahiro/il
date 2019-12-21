package logger

import "time"

import "os"

func ExampleDebugf() {
	std = new(os.Stdout)
	now = func() time.Time { return time.Date(2019, 12, 31, 23, 59, 59, 0, time.UTC) }

	Debugf("%s", "test")

	// Output:
	// 2019/12/31 23:59:59 test
}
