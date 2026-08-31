package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TCPServer `yaml:"tcp_server"`
	DataBase  `yaml:"data_base"`
}

type TCPServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type DataBase struct {
	Path string `yaml:"path"`
}

func MustLoad() *Config {
	var cfg Config

	ConfigPath := os.Getenv("CONFIG_PATH")
	if ConfigPath == "" {
		log.Fatal("CONFIG_PATH + environment variable not set")
	}
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatal("read config failed")
	}
	return &cfg
}
