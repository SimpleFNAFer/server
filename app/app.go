package app

import (
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/simplefnafer/network-attack-lab/server/app/home/handler"
	"github.com/simplefnafer/network-attack-lab/server/app/home/repository"
	"github.com/simplefnafer/network-attack-lab/server/app/home/usecase"
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func Run() {
	ConfigureLogger()

	db := RunDB()
	rdb := ConnectRedis()
	dur := LoadSessionDuration()

	r := repository.NewRepository(db, rdb, dur)

	uc := usecase.NewUseCase(r)

	h := handler.NewHandler(uc, dur)

	routes := ConfigureRoutes(h)

	metrics := ConfigureMetrics()

	go uc.UpdateCountLastSecondRequests(metrics)

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

func ConnectRedis() *redis.Client {
	redisUser := os.Getenv("REDIS_USER")
	redisPassword := os.Getenv("REDIS_USER_PASSWORD")
	host := os.Getenv("REDIS_HOST")

	url := fmt.Sprintf(
		"redis://%s:%s@%s:6379/0?protocol=3",
		redisUser,
		redisPassword,
		host)
	opts, err := redis.ParseURL(url)
	if err != nil {
		logrus.WithField("origin.function", "ConnectRedis").Error(err)
	}

	return redis.NewClient(opts)
}

func LoadSessionDuration() time.Duration {
	durStr := os.Getenv("SESSION_DURATION")
	dur, err := time.ParseDuration(durStr)
	if err != nil {
		logrus.WithField("origin.function", "LoadSessionDuration").Error(err)
	}

	return dur
}

func ConfigureRoutes(h home.Handler) (routes *chi.Mux) {
	routes = chi.NewRouter()

	routes.Handle(
		"/{style}.css",
		http.FileServer(
			http.Dir("template")))

	routes.Handle(
		"/{script}.js",
		http.FileServer(
			http.Dir("template")))

	routes.Get("/", h.Home)
	routes.Get("/admin", h.Admin)
	routes.Get("/admin-login", h.AdminLoginGet)
	routes.Get("/forbidden", h.Forbidden)
	routes.Post("/admin-login", h.AdminLoginPost)
	routes.Post("/block", h.Block)
	routes.Post("/unblock", h.Unblock)

	return routes
}

func ConfigureMetrics() *model.Metrics {
	metrics := model.Metrics{
		LastSecondRequests: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "last_second_requests",
			},
			[]string{"ip"},
		),
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

	return &metrics
}
