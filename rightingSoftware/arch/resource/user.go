package resource

import "time"

type User struct {
	ID        int64
	Name      string
	Age       int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
