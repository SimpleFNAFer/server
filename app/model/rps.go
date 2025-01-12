package model

type RPS struct {
	SourceIP string `db:"source_ip"`
	Total    int    `db:"total"`
}
