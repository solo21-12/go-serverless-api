package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	AWS_REGION string `mapstructure:"AWS_REGION"`
	TABLE_NAME string `mapstructure:"TABLE_NAME"`
}

func NewEnv() *Env {
	projectRoot, err := filepath.Abs(filepath.Join(""))

	if err != nil {
		log.Fatalf("Error getting project root: %v", err)
	}

	viper.SetConfigFile(filepath.Join(projectRoot, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	env := &Env{}

	if err := viper.Unmarshal(env); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return env
}
