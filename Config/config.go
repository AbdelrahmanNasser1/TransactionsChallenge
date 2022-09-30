package Config

import (
	"errors"
	"github.com/spf13/viper"
)

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	ConnectionString string
}

type KafkaConfiguration struct {
	URL   string
	Topic string
}

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Kafka    KafkaConfiguration
}

func SetUpViper(configurations Configurations) (Configurations, error) {
	viper.SetConfigName("../config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return configurations, errors.New("failed to read config file")
	}
	err := viper.Unmarshal(&configurations)
	if err != nil {
		return configurations, errors.New("failed to Convert to object")
	}
	return configurations, nil
}
