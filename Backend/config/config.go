package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	AppConfig  Config
	configOnce sync.Once
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

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

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	JwtSecret string `yaml:"jwtSecret"`

	IpfsApiKey string `mapstructure:"ipfsApiKey"`
}

// LoadConfig loads the configuration from the YAML file.
func LoadConfig(configPath string) error {
	var err error
	configOnce.Do(func() {
		var data []byte
		data, err = os.ReadFile(configPath)
		if err != nil {
			log.Printf("Error reading config file: %v", err)
			return
		}
		err = yaml.Unmarshal(data, &AppConfig)
		if err != nil {
			log.Printf("Error unmarshalling config file: %v", err)
			return
		}
	})
	return err
}
