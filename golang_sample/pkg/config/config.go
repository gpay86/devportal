package config

import (
	"github.com/spf13/viper"
)

// LoadConfig from file
func LoadConfig(configPath, configName string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName(configName + ".json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}

	// set project root directory to config

	return result, nil
}
