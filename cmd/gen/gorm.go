package main

import (
	viper2 "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	viper2.SetConfigFile("config.yml")
	err := viper2.ReadInConfig()
	if err != nil {
		return
	}
	open, err := gorm.Open(mysql.Open(viper2.GetString("Dsn")))
	if err != nil {
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "dao",
		ModelPkgPath: "model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	})

	g.UseDB(open)
	g.ApplyBasic(g.GenerateAllTable()...)
	//g.ApplyBasic(g.GenerateModel("user"))
	g.Execute()
}
