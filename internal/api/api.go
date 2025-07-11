package api

import "github.com/chiquitav2/journalful/pkg/conf"

type ApiModule interface {
	Register() error
	Start(config *conf.Config) error
}
