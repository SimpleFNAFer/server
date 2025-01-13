package home

import (
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"html/template"
)

type UseCase interface {
	Home(sourceIP string) (tmpl *template.Template, data model.HomeData)
	Admin() (tmpl *template.Template, data model.AdminData)
	AdminLoginTemplate() (tmpl *template.Template)
	Forbidden() (tmpl *template.Template)
	ForbiddenUser() (tmpl *template.Template)

	AdminLogin(login, password string) string
	AdminCheckSession(session string) bool
	BlockIP(ip string)
	UnblockIP(ip string)
	IsBlockedIP(ip string) bool

	UpdateCountLastSecondRequests(m *model.Metrics)
}
