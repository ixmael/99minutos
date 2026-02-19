package main

// ApplicationConfig represents the configuration for the application.
type ApplicationConfig struct {
	Environment string `toml:"environment"`
	RestAPI     struct {
		Port int `toml:"port"`
	} `toml:"restapi"`
	Repository struct {
		PostgresURL string `toml:"postgres"`
	} `toml:"repository"`
	Queue struct {
		RabbitMQURL string `toml:"rabbitmq"`
	} `toml:"queue"`
}
