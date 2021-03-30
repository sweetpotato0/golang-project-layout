package entity

import (
	uuid "github.com/iris-contrib/go.uuid"
)

var _ Entity = (*entity)(nil)

//Entity is the entity's father interface.
type Entity interface {
	Identity() string
}

type entity struct {
	entityName   string
	identity     string
	entityObject interface{}
}

func (e *entity) Identity() string {
	if e.identity == "" {
		u, _ := uuid.NewV1()
		e.identity = u.String()
	}
	return e.identity
}
