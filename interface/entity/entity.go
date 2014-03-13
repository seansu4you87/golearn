package main

import (
	"net/http"
)

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, ent Entity) error {
	return ent.UnmarshalHTTP(r)
}

type User struct {
}

func (u *User) UnmarshalHTTP(r *http.Request) error {
	return nil
}
