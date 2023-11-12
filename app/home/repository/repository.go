package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/sirupsen/logrus"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) home.Repository {
	return &repository{db: db}
}

func (r *repository) SaveRequest(sourceIP string) {
	if _, err := r.db.Exec(querySaveRequest, sourceIP); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.SaveRequest",
		).Error(err)
	}
}

func (r *repository) CountRequests() (count int) {
	count = 0

	if err := r.db.Get(&count, queryCountRequests); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.CountRequests",
		).Error(err)
	}

	return
}

func (r *repository) CountLastSecondRequests() (count int) {
	count = 0

	if err := r.db.Get(&count, queryCountLastSecondRequest); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.CountLastSecondRequests",
		).Error(err)
	}

	return count
}
