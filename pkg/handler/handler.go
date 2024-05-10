package handler

import "github.com/Pineapple217/Netlane/pkg/ent"

type Handler struct {
	Dbc *ent.Client
}

func NewHandler(dbc *ent.Client) *Handler {
	return &Handler{
		Dbc: dbc,
	}
}
