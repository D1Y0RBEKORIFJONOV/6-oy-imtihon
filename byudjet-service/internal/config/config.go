package config

import "os"

type Config struct {
	Budget struct{
		Host string
		Port string
	}
	Mongo struct {
		Url string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Budget.Host=osGetenv("BUDGET_HOST","tcp")
	c.Budget.Port = osGetenv("BUDGET_PORT", "byudjet-service:8888")

	c.Mongo.Url=osGetenv("MONGO_URL","mongodb://mongo:27017")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
