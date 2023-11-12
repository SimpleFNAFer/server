package home

import (
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"html/template"
)

type UseCase interface {
	Home(sourceIP string) (tmpl *template.Template, data model.HomeData)
	UpdateCountLastSecondRequests() float64
}
