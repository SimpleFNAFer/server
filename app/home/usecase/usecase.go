package usecase

import (
	"github.com/google/uuid"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"slices"
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

	data = model.HomeData{
		Total: total,
	}

	return tmpl, data
}

func (uc *useCase) UpdateCountLastSecondRequests(m *model.Metrics) {
	for {
		m.Lock()

		rps := uc.r.CountLastSecondRequests()

		for _, v := range rps {
			m.LastSecondRequests.WithLabelValues(v.SourceIP).Set(float64(v.Total))
		}

		m.Unlock()
	}
}

func (uc *useCase) AdminLogin(login, password string) string {
	passwordHash := uc.r.GetAdminPassword(login)
	if passwordHash == "" {
		return ""
	}

	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) != nil {
		return ""
	}

	session := uuid.NewString()

	if err := uc.r.SaveSession(session); err != nil {
		return ""
	}

	return session
}

func (uc *useCase) AdminCheckSession(session string) bool {
	exists, err := uc.r.CheckSession(session)
	if err != nil {
		return false
	}

	return exists
}

func (uc *useCase) Admin() (tmpl *template.Template, data model.AdminData) {
	tmpl, err := template.ParseFiles("template/admin.html")
	if err != nil {
		logrus.WithField("origin.function", "PrepareTemplate").Error(err)
	}

	freqIPs := uc.r.GetFrequentIPs()
	blockedIPs := uc.r.GetBlockedIPs()

	data = model.AdminData{
		FrequentIPs: freqIPs,
		BlockedIPs:  blockedIPs,
	}

	return tmpl, data
}

func (uc *useCase) AdminLoginTemplate() (tmpl *template.Template) {
	tmpl, err := template.ParseFiles("template/admin-login.html")
	if err != nil {
		logrus.WithField("origin.function", "PrepareTemplate").Error(err)
	}

	return tmpl
}

func (uc *useCase) Forbidden() (tmpl *template.Template) {
	tmpl, err := template.ParseFiles("template/forbidden.html")
	if err != nil {
		logrus.WithField("origin.function", "PrepareTemplate").Error(err)
	}

	return tmpl
}

func (uc *useCase) BlockIP(ip string) {
	uc.r.BlockIP(ip)
}

func (uc *useCase) UnblockIP(ip string) {
	uc.r.UnblockIP(ip)
}

func (uc *useCase) IsBlockedIP(ip string) bool {
	blocked := uc.r.GetBlockedIPs()
	if slices.Contains(blocked, ip) {
		return true
	}

	return false
}
