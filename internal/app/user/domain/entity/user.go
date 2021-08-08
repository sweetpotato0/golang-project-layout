package entity

import (
	"fmt"
	"strconv"

	"layout/internal/app/user/domain/po"
	"layout/internal/pkg/entity"
)

// User represent entity of the user
type User struct {
	entity.Entity
	po.User
}

// Identity .
func (a *User) Identity() string {
	return strconv.FormatInt(a.ID, 10)
}

// NewUser initialize User
func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name")
	}

	a := &User{User: po.User{Name: name}}
	return a, nil
}
