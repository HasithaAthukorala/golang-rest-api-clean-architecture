package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var lock = &sync.Mutex{}
var config Config

type Config struct {
	DbDSN                       string `json:"dbDSN"`
	LocationVerificationHostURL string `json:"locationVerificationHostURL"`
	ServiceBusConnectionString  string `json:"serviceBusConnectionString"`
}

func Load(fileName string) (*Config, error) {
	// make this thread safe
	lock.Lock()
	defer lock.Unlock()

	// check if empty and create
	if (config == Config{}) {
		config = Config{}

		viper.SetConfigName(fileName)
		viper.SetConfigType("json")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error loading application config file :%v", err)
		}
		if err := viper.Unmarshal(&config); err != nil {
			return nil, fmt.Errorf("malformed JSON config file :%v", err)
		}
	}

	return &config, nil
}
