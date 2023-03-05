package config

type SecConfig interface {
	GetJWTSecret() string
}
