package main

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	notFound = errors.New("user not found")
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Key       string `json:"key"`
	Sent      int    `json:"sent"`
	Confirmed bool   `json:"confirmed"`
}

type requestResponse struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type requestJSON struct {
	Email   string `json:"email"`
	Version string `json:"version"`
}

type errorResponse struct {
	Error string `json:"error"`
}
