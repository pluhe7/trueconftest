package handler

import (
	"github.com/pluhe7/trueconftest/model"
)

type Handler struct {
	Users map[int]*model.User
	Seq   int
}

func NewHandler(us map[int]*model.User, seq int) *Handler {
	return &Handler{
		Users: us,
		Seq:   seq,
	}
}
