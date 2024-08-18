package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env:"ENV" env-default:"local"`
	Storage Storage  `yaml:"database"`
}

type Storage struct {
	DbHost     string `yaml:"host" env-default:"localhost"`
	DbPort     string `yaml:"port" env-default:"5432"`
	DbUser     string `yaml:"user"`
	DbPassword string `yaml:"password"`
	DbName     string `yaml:"name"`
}

type HTTPserver struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"6s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func LoadConfig() (*Config, error) {
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatalf("config_file %s does not exist", config_path)
	}

	var cfg Config

	err := cleanenv.ReadConfig(config_path, &cfg)
	if err != nil {
		log.Fatalf("can not read config %s", err)
	}
    

	return &cfg, nil

}
