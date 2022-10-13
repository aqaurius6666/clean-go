package config

type AppConfig struct {
	Version string `cli:"version" env:"VERSION" default:"0.1.0"`
	Log     LogConfig
	Http    HTTPConfig
	Db      DBConfig
	Auth    AuthConfig
	Es      ESConfig
	Event   EventConfig
	Otel    OTELConfig
}

type LogConfig struct {
	Level  string `cli:"log-level" env:"LEVEL" default:"info"`
	Format string `cli:"log-format" env:"FORMAT" default:"text"`
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

type OTELConfig struct {
	Enabled       bool   `cli:"otel-enabled" env:"ENABLED" default:"false"`
	CollectorAddr string `cli:"otel-collector-addr" env:"COLLECTOR_ADDR" default:"localhost:4317"`
	ID            int64  `cli:"otel-id" env:"ID" default:"1"`
	ServiceName   string `cli:"otel-service-name" env:"SERVICE_NAME" default:"clean-go"`
}

type AuthConfig struct {
	Secret                string `cli:"auth-secret" env:"SECRET" default:"thisissecretforjwt"`
	ExpireDuration        int64  `cli:"auth-expire-duration" env:"EXPIRE_DURATION" default:"1800"`
	RefreshExpireDuration int64  `cli:"auth-refresh-expire-duration" env:"REFRESH_EXPIRE_DURATION" default:"864000"`
}

type ESConfig struct {
	DSN    string `cli:"es-dsn" env:"DSN" default:""`
	Host   string `cli:"es-host" env:"HOST" default:"localhost"`
	User   string `cli:"es-user" env:"USER" default:"admin"`
	Pass   string `cli:"es-pass" env:"PASS" default:"changeit"`
	Port   string `cli:"es-port" env:"PORT" default:"2113"`
	Scheme string `cli:"es-scheme" env:"SCHEME" default:"esdb"`
}

type EventConfig struct {
	StreamRegex string `cli:"event-stream-regex" env:"STREAM_REGEX" default:"^event-"`
}
