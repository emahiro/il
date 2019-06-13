package model

import "log"

// UserA ...
type UserA struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// GetUserA ...
func GetUserA() (*UserA, error) {
	return &UserA{Name: "ema", Age: 29}, nil
}

// User ...
type User interface {
	Get() (*UserB, error)
}

// UserB ...
type UserB struct {
	Name string
	Age  int64
}

// Get ...
func (u *UserB) Get() (*UserB, error) {
	log.Printf("implementation detail of UserB. This UserName is %s", u.Name)
	return u, nil
}
