package config

type Config struct {
	Version string `default:""`
	Port    string `default:"8000"`

	DBHost     string `envconfig:"db_host" default:"localhost"`
	DBPort     string `envconfig:"db_port" default:"3306"`
	DBName     string `envconfig:"db_name" default:"pear_system"`
	DBUser     string `envconfig:"db_user" default:"root"`
	DBPassword string `envconfig:"db_password" default:"root"`
}
