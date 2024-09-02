package Config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfStruct struct {
	Mysql  string `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	JwtKey string `yaml:"jwtKey"`
}
type Redis struct {
	Url string `yaml:"url"`
	Pwd string `yaml:"pwd"`
	DB  int    `yaml:"DB"`
}

var Conf = getConf()

func getConf() ConfStruct {
	var conf ConfStruct
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		return conf
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return conf
	}
	return conf
}
