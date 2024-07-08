package handler

import (
	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
)

type Handler struct {
	Auth pb.AuthServiceClient
}

func NewHandler(auth pb.AuthServiceClient) *Handler {
	return &Handler{
		Auth: auth,
	}
}
