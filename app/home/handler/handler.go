package handler

import (
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	uc home.UseCase
}

func NewHandler(uc home.UseCase) home.Handler {
	return &handler{uc: uc}
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	sourceIP := r.RemoteAddr

	template, data := h.uc.Home(sourceIP)
	err := template.Execute(w, data)
	if err != nil {
		logrus.WithField("origin.function", "Home").Error(err)
	}
}
