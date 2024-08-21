package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		DBUser     string `yaml:"dbuser"`
		DBPassword string `yaml:"dbpassword"`
		DBHost     string `yaml:"dbhost"`
		DBPort     int    `yaml:"dbport"`
		DBName     string `yaml:"dbname"`
	} `yaml:"mysql"`
}

func LoadConfig() Config {
	// 读取YAML文件
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 解析YAML
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
