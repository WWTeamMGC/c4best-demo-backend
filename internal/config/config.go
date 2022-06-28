package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

type Config struct {
	Debug bool  `yaml:"debug"`
	Http  Http  `yaml:"http"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Kafka Kafka `yaml:"kafka"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database int    `yaml:"database"`
	Password string `yaml:"password"`
	Enable   int    `yaml:"enable"`
}
type Kafka struct {
	Enable  bool     `yaml:"enable"`
	Brokers []string `yaml:"brokers"`
}

var (
	cfg  *Config
	once sync.Once
)

func Phase() (*Config, error) {
	once.Do(func() {
		cfg = &Config{}
		configfile, err := os.ReadFile("./config.yaml")
		if err != nil {
			panic(err)
		}
		if len(configfile) == 0 {
			panic("config file not found!")
		}
		err = yaml.Unmarshal([]byte(configfile), cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg, nil
}
