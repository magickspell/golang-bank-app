package config

import (
	logg "backend/logger"
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	Host  string
	DbURL string
}

func GetConfig(log *logg.Logger) *Config {
	dbURL := os.Getenv("GO_DB_URL")
	host := os.Getenv("GO_HOST")
	cfg := &Config{
		Host:  host,
		DbURL: dbURL,
	}

	keyValues := reflect.ValueOf(cfg)
	for i := 0; i < keyValues.NumField(); i++ {
		field := keyValues.Field(i)
		fmt.Printf("Field Name: %v", field)
		// fmt.Printf("Field Name: %s, Type: %s, Tag: '%s'\n", field.Name, field.Type, field.Tag)
	}
	if len(dbURL) == 0 || len(host) == 0 {
		log.OuteputLog(logg.LogPayload{Error: fmt.Errorf("env has empty value")})
	}
	return cfg
}
