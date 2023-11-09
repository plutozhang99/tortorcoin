package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for our application. Each field corresponds to a section in the config file.
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

// DatabaseConfig holds the DSN (Data Source Name) for the database connection.
type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

// JWTConfig holds configuration related to JWT, such as the secret key used for signing tokens.
type JWTConfig struct {
	SecretKey      string `mapstructure:"secret_key"`
	ExpirationTime int64  `mapstructure:"expiration_time"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("pkg/config") // path to look for the config file in
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name

	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return
	}

	// Unmarshal the config file into the Config struct.
	err = viper.Unmarshal(&config)
	return
}
