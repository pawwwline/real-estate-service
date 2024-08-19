package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-default:"local"`
	Storage    Storage    `yaml:"db"`
	HTTPserver HTTPserver `yaml:"httpserver"`
}

type Storage struct {
	DbHost     string `yaml:"host" env-default:"localhost"`
	DbPort     string `yaml:"port" env-default:"5432"`
	DbUser     string `yaml:"user" env-default:"myuser"`
	DbPassword string `yaml:"password" env-default:"pass"`
	DbName     string `yaml:"name" env-default:"estatedb"`
}

type HTTPserver struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"6s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config_file %s does not exist", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config %s", err)
	}

	return &cfg, nil

}
