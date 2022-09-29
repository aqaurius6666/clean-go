package config

type AppConfig struct {
	Version string `cli:"version" env:"VERSION" default:"0.1.0"`
	Log     LogConfig
	HTTP    HTTPConfig
	Db      DBConfig
}

type LogConfig struct {
	Level string `cli:"log-level" env:"LEVEL" default:"info"`
}

type HTTPConfig struct {
	Port string `cli:"http-port" env:"PORT" default:"80"`
}
type DBConfig struct {
	DSN    string `cli:"db-dsn" env:"DSN" default:""`
	Host   string `cli:"db-host" env:"HOST" default:"localhost"`
	Port   string `cli:"db-port" env:"PORT" default:"5432"`
	User   string `cli:"db-user" env:"USER" default:"postgres"`
	Pass   string `cli:"db-pass" env:"PASS" default:"postgres"`
	Name   string `cli:"db-name" env:"NAME" default:"postgres"`
	Scheme string `cli:"db-scheme" env:"SCHEME" default:"postgres"`
	Query  string `cli:"db-query" env:"QUERY" default:"sslmode=disable"`
}
