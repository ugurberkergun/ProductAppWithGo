package app

import "ProductAppWithGo/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	config := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: config,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                   "localhost",
		Port:                   "6432",
		DbName:                 "productapp",
		UserName:               "postgres",
		Password:               "postgres",
		MaxConnections:         "10",
		MaxConnectionsIdleTime: "30s",
	}
}
