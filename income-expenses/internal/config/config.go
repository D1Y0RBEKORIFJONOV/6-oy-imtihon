package config

import "os"

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBname   string
	}
	Budget struct {
		Port string
	}
	Notification struct {
		Port string
	}
	IncomeExpenses struct {
		Host string
		Port string
	}
	Kafka struct {
		Port  string
		Topic string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Database.User = osGetenv("DB_USER", "postgres")
	c.Database.Password = osGetenv("DB_PASSWORD", "+_+2005+_+")
	c.Database.Host = osGetenv("DB_HOST", "postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "postgres")

	c.Budget.Port = osGetenv("BUDGET_PORT", "byudjet-service:8888")

	c.Notification.Port = osGetenv("NOTIFICATION_PORT", "notification_service:9001")

	c.IncomeExpenses.Host = osGetenv("INCOMEEXPENSES_HOST", "tcp")
	c.IncomeExpenses.Port = osGetenv("INCOMEEXPENSES_PORT", "income-expenses_container:8080")

	c.Kafka.Port = osGetenv("KAFKA_PORT", "broker:29092")
	c.Kafka.Topic = osGetenv("KAFKA_TOPIC", "incomeexpenses17")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
