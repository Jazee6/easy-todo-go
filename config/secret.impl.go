package config

import viper2 "github.com/spf13/viper"

type CfgS struct {
}

var vs = viper2.New()

func init() {
	vs.SetConfigFile("secret.yml")
	err := vs.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func NewCfgS() *CfgS {
	return &CfgS{}
}

func (c CfgS) GetJWTSecret() string {
	return vs.GetString("JWTSecret")
}
