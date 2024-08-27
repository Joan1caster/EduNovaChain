package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	config     *Config
	configOnce sync.Once
)

type Config struct {
	MySQL struct {
		DBUser     string `yaml:"dbuser"`
		DBPassword string `yaml:"dbpassword"`
		DBHost     string `yaml:"dbhost"`
		DBPort     int    `yaml:"dbport"`
		DBName     string `yaml:"dbname"`
	} `yaml:"mysql"`

	Contract struct {
		Eth_rpc_url     string `yaml:"eth_rpc_url"`
		UsrAddress      string `yaml:"usrAddress"`
		PrivateKey      string `yaml:"privateKey"`
		ContractAddress string `yaml:"contractAddress"`
		ContractApiFile string `yaml:"contractApiFile"`
	} `yaml:"contract"`
}

func LoadConfig() *Config {
	configOnce.Do(func() {
		// 读取YAML文件
		data, err := os.ReadFile("config/config.yaml")
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// 解析YAML
		err = yaml.Unmarshal(data, config)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	})
	return config
}
