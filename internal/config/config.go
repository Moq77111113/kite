package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)


type Config struct {
	Registry string `mapstructure:"registry"`
	Out      string `mapstructure:"out"`
	Flavor   string `mapstructure:"flavor"`
}

var configDir string

func init() {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".kite")

	viper.AddConfigPath(configDir)
	viper.SetConfigName("kite")
	viper.SetConfigType("yaml")

	viper.SetDefault("registry", "./registry.json")
	viper.SetDefault("out", "./")
	viper.SetDefault("flavor", "default")

}



func Load() error {
	return viper.ReadInConfig()
}

func Save() error {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			return err
		}
	}
	return viper.WriteConfigAs(filepath.Join(configDir, ".kite"))
}

func InitConfig() error {
	return Save()
}

func GetConfig() (*Config, error) {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {

		return nil, err
	}
	return &config, nil
}

