package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Adress      string        `yaml:"port" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

func LoadConfig() *Config {
	bytes, err := os.ReadFile("./config-path.txt")
	if err != nil {
		log.Fatal(err)
	}
	configPath := string(bytes)
	//check if file exists
	if configPath == "" {
		log.Fatalf("Config file: %s, is empty or broken", configPath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", configPath)
	}

	return &cfg
}
