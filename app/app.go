package app

import (
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/simplefnafer/network-attack-lab/server/app/home/handler"
	"github.com/simplefnafer/network-attack-lab/server/app/home/repository"
	"github.com/simplefnafer/network-attack-lab/server/app/home/usecase"
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var metrics model.Metrics

func Run() {
	ConfigureLogger()

	db := RunDB()
	r := repository.NewRepository(db)

	uc := usecase.NewUseCase(r)

	h := handler.NewHandler(uc)

	routes := ConfigureRoutes(h)

	ConfigureMetrics(uc)

	logrus.WithField("origin.function", "Run").Info("server started")
	if err := http.ListenAndServe(":8080", routes); err != nil {
		logrus.WithField("origin.function", "Run").Error(err)
	}
}

func ConfigureLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "02.01.2006 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatter.DisableLevelTruncation = true
	logrus.SetFormatter(customFormatter)
}

func RunDB() (db *sqlx.DB) {
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")

	dataSource := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		dbUsername,
		dbPassword,
		host,
		dbName)
	db, err := sqlx.Connect("pgx", dataSource)
	if err != nil {
		logrus.WithField("origin.function", "RunDB").Error(err)
	}

	return
}

func ConfigureRoutes(h home.Handler) (routes *chi.Mux) {
	routes = chi.NewRouter()

	routes.Handle(
		"/{style}.css",
		http.FileServer(
			http.Dir("template")))

	routes.Get("/", h.Home)

	return routes
}

func ConfigureMetrics(uc home.UseCase) {
	metrics = model.Metrics{
		LastSecondRequests: prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "last_second_requests",
			},
			uc.UpdateCountLastSecondRequests),
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(metrics.LastSecondRequests)

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

		if err := http.ListenAndServe(":8082", mux); err != nil {
			logrus.WithField(
				"origin.function",
				"ConfigureMetrics",
			).Error(err)
		}
	}()
}
