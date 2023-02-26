package config

type Config interface {
	GetWhiteList() []string
	GetVersion() string
}
