package config

import (
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var (
	//go:embed secrets/id_rsa_mock
	secretKey embed.FS
)

func NewConfig() error {
	viper.SetConfigName("config")             // name of config file (without extension)
	viper.SetConfigType("yaml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/mock_server/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.mock_server") // call multiple times to add many search paths
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}

	return nil
}

type App struct {
	Addr          string `mapstructure:"addr"`
	SecretKeyPath string `mapstructure:"secret_path"`
}

func ParseSecret(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		data, err := secretKey.ReadFile("secrets/id_rsa_mock")
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}
