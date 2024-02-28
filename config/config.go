package config

import "github.com/spf13/viper"

type Config struct {
	TeachableURL    string
	TeachableAPIKey string
}

func LoadEnvVars() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	return &Config{
		TeachableURL:    viper.GetString("TEACHBLE_URL"),
		TeachableAPIKey: viper.GetString("TEACHABLE_API_KEY"),
	}
}
