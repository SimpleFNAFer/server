package home

import "github.com/simplefnafer/network-attack-lab/server/app/model"

type Repository interface {
	SaveRequest(sourceIP string)
	CountRequests() (count int)
	CountLastSecondRequests() []model.RPS
	GetFrequentIPs() []string
	GetBlockedIPs() []string
	BlockIP(ip string)
	UnblockIP(ip string)
	GetAdminPassword(login string) string

	SaveSession(s string) error
	CheckSession(s string) (bool, error)
}
