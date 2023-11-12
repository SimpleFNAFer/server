package usecase

import (
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"github.com/sirupsen/logrus"
	"html/template"
)

type useCase struct {
	r home.Repository
}

func NewUseCase(r home.Repository) home.UseCase {
	return &useCase{r: r}
}

func (uc *useCase) Home(sourceIP string) (tmpl *template.Template, data model.HomeData) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		logrus.WithField("origin.function", "PrepareTemplate").Error(err)
	}

	uc.r.SaveRequest(sourceIP)

	total := uc.r.CountRequests()

	data = model.HomeData{Total: total}

	return tmpl, data
}

func (uc *useCase) UpdateCountLastSecondRequests() float64 {
	return float64(uc.r.CountLastSecondRequests())
}
