package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/simplefnafer/network-attack-lab/server/app/home"
	"github.com/simplefnafer/network-attack-lab/server/app/model"
	"github.com/sirupsen/logrus"
	"time"
)

type repository struct {
	db              *sqlx.DB
	rdb             *redis.Client
	sessionDuration time.Duration
}

func NewRepository(
	db *sqlx.DB,
	rdb *redis.Client,
	sessionDuration time.Duration,
) home.Repository {
	return &repository{
		db:              db,
		rdb:             rdb,
		sessionDuration: sessionDuration,
	}
}

func (r *repository) SaveRequest(sourceIP string) {
	if _, err := r.db.Exec(querySaveRequest, sourceIP); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.SaveRequest",
		).Error(err)
	}
}

func (r *repository) GetAdminPassword(login string) string {
	passwordHash := ""

	if err := r.db.Get(&passwordHash, queryGetAdminPassword, login); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.GetAdminPassword",
		).Error(err)
	}

	return passwordHash
}

func (r *repository) SaveSession(s string) error {
	return r.rdb.Set(context.Background(), s, struct{}{}, r.sessionDuration).Err()
}

func (r *repository) CheckSession(s string) (bool, error) {
	err := r.rdb.Get(context.Background(), s).Err()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
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

func (r *repository) CountLastSecondRequests() []model.RPS {
	res := make([]model.RPS, 0)

	if err := r.db.Select(&res, queryCountLastSecondRequest); err != nil {
		logrus.WithField(
			"origin.function",
			"repository.CountLastSecondRequests",
		).Error(err)
	}

	return res
}

func (r *repository) GetFrequentIPs() []string {
	res := make([]string, 0)

	if err := r.db.Select(&res, queryGetFrequentIPs); err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.WithField(
			"origin.function",
			"repository.GetFrequentIPs",
		).Error(err)
	}

	return res
}

func (r *repository) GetBlockedIPs() []string {
	res := make([]string, 0)

	if err := r.db.Select(&res, queryGetBlockedIPs); err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.WithField(
			"origin.function",
			"repository.GetBlockedIPs",
		).Error(err)
	}

	return res
}

func (r *repository) BlockIP(ip string) {
	_, err := r.db.Exec(queryBlockIP, ip)
	if err != nil {
		logrus.WithField(
			"origin.function",
			"repository.BlockIP",
		).Error(err)
	}
}

func (r *repository) UnblockIP(ip string) {
	_, err := r.db.Exec(queryUnblockIP, ip)
	if err != nil {
		logrus.WithField(
			"origin.function",
			"repository.UnblockIP",
		).Error(err)
	}
}
