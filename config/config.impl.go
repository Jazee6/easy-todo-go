package config

import (
	viper2 "github.com/spf13/viper"
)

type Cfg struct {
}

func NewCfg() *Cfg {
	return &Cfg{}
}

var v = viper2.New()

func init() {
	v.SetConfigFile("config.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func (c Cfg) GetWhiteList() []string {
	return v.GetStringSlice("WhiteList")
}

func (c Cfg) GetVersion() string {
	return v.GetString("Version")
}

func (c Cfg) GetSecret() string {
	//TODO implement me
	panic("implement me")
}
