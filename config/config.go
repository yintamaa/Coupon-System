package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {

	// repository
	MysqlDSN string `yaml:"mysql_dsn"`
}

var config *Config

func GetCouponConfig() *Config {
	return config
}

func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	SetConfig(c)
	logrus.Infof("NewConfig: %#v", c)
	return c, nil
}

func SetConfig(c *Config) {
	logrus.Infof("SetConfig: %#v", c)
	config = c
}
