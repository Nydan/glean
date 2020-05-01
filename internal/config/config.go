package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type (
	// Config struct to wrap config file
	Config struct {
		HTTPServer HTTPServer
		Logger     Logger
		Database   Database `mapstructure:"database"`
		Redis      Redis

		Version   string `yaml:"-"`
		BuildDate string `yaml:"-"`
	}

	// HTTPServer config for server setting
	HTTPServer struct {
		ListenAddress   string
		Port            int
		GracefulTimeout time.Duration
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		IdleTimeout     time.Duration
	}

	// Logger logging setting
	Logger struct {
		EnableConsole     bool
		ConsoleJSONFormat bool
		ConsoleLevel      string
		EnableFile        bool
		FileJSONFormat    bool
		FileLevel         string
		FileLocation      string
	}

	// Database configurations
	Database struct {
		Master  string
		Replica string
	}

	// Redis configurations
	Redis struct {
		Endpoint string
		Timeout  int
		MaxIdle  int
	}
)

var config Config

// Load loads the config file into Config struct
func Load(env string) (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// More option of config path can be added here
	viper.AddConfigPath(fmt.Sprintf("./config/"))                       // Staging or Production
	viper.AddConfigPath(fmt.Sprintf("files/etc/config/%s/", env))       // Unix Local
	viper.AddConfigPath(fmt.Sprintf("../../files/etc/config/%s/", env)) // Windows Local

	// Get the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Convert into struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
