package api

type ApiModule interface {
	Register() error
	Start() error
}
