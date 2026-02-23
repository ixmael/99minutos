package main

// ApplicationConfig represents the configuration for the application.
type ApplicationConfig struct {
	Environment string `toml:"environment"`
	Version     string `toml:"version"`
	RestAPI     struct {
		Port int `toml:"port"`
	} `toml:"restapi"`
	Repository struct {
		PostgresURL            string `toml:"postgres"`
		PostgresMigrationsPath string `toml:"migrations_path"`
	} `toml:"repository"`
	Queue struct {
		RabbitMQURL string `toml:"rabbitmq"`
	} `toml:"queue"`
	Cache struct {
		ValkeyURL string `toml:"valkey"`
	} `toml:"cache"`
}
