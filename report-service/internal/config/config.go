package config

import "os"

type Config struct {
	Report struct{
		Host string
		Port string
	}
	Budget struct {
		Port string
	}
	IncomeExpenses struct {
		Port string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Budget.Port = osGetenv("BUDGET_PORT", "byudjet-service:8888")

	c.IncomeExpenses.Port = osGetenv("INCOMEEXPENSES_PORT", "income-expenses_container:8080")

	c.Report.Host=osGetenv("REPORT_HOST","tcp")
	c.Report.Port=osGetenv("REPORT_PORT","report-service:8000")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
