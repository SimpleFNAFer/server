package handler

import (
	"errors"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type handler struct {
	uc              home.UseCase
	sessionDuration time.Duration
}

func NewHandler(
	uc home.UseCase,
	sessionDuration time.Duration,
) home.Handler {
	return &handler{
		uc:              uc,
		sessionDuration: sessionDuration,
	}
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	sourceIP := strings.Split(r.RemoteAddr, ":")[0]

	if h.uc.IsBlockedIP(sourceIP) {
		http.Redirect(w, r, "/forbidden-user", http.StatusSeeOther)
		return
	}

	template, data := h.uc.Home(sourceIP)
	err := template.Execute(w, data)
	if err != nil {
		logrus.WithField("origin.function", "Home").Error(err)
	}
}

func (h *handler) Admin(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	exists := h.uc.AdminCheckSession(c.Value)
	if !exists {
		http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
		return
	}

	template, data := h.uc.Admin()
	err = template.Execute(w, data)
	if err != nil {
		logrus.WithField("origin.function", "Admin").Error(err)
	}
}

func (h *handler) AdminLoginGet(w http.ResponseWriter, _ *http.Request) {
	template := h.uc.AdminLoginTemplate()
	err := template.Execute(w, nil)
	if err != nil {
		logrus.WithField("origin.function", "AdminLoginTemplate").Error(err)
	}
}

func (h *handler) AdminLoginPost(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	session := h.uc.AdminLogin(login, password)

	if session == "" {
		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	cookie := &http.Cookie{
		Name:    "session",
		Value:   session,
		Path:    "/",
		Expires: time.Now().Add(h.sessionDuration),
	}

	http.SetCookie(w, cookie)
	r.AddCookie(cookie)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *handler) Block(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	exists := h.uc.AdminCheckSession(c.Value)
	if !exists {
		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	ip := r.FormValue("ip")
	h.uc.BlockIP(ip)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *handler) Unblock(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	exists := h.uc.AdminCheckSession(c.Value)
	if !exists {
		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
		return
	}

	ip := r.FormValue("ip")
	h.uc.UnblockIP(ip)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *handler) Forbidden(w http.ResponseWriter, _ *http.Request) {
	template := h.uc.Forbidden()
	err := template.Execute(w, nil)
	if err != nil {
		logrus.WithField("origin.function", "Forbidden").Error(err)
	}
}

func (h *handler) ForbiddenUser(w http.ResponseWriter, _ *http.Request) {
	template := h.uc.ForbiddenUser()
	err := template.Execute(w, nil)
	if err != nil {
		logrus.WithField("origin.function", "ForbiddenUser").Error(err)
	}
}
