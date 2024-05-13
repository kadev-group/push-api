package models

type Config struct {
	ENV        string `env:"APP_ENV"`
	ServerPORT string `env:"SERVER_PORT"`

	PSQL
	Redis
	SMTP
	RabbitMQ
}

func (c Config) Env() string {
	return c.ENV
}

type PSQL struct {
	PqHOST     string `env:"PSQL_HOST"`
	PqPORT     string `env:"PSQL_PORT"`
	PqUSER     string `env:"PSQL_USER"`
	PqPASSWORD string `env:"PSQL_PASSWORD"`
	PqDATABASE string `env:"PSQL_DATABASE"`
	PqSSL      string `env:"PSQL_SSL"`
}

type Redis struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDatabase int    `env:"REDIS_DATABASE"`
}

type SMTP struct {
	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     int    `env:"SMTP_PORT"`
	SmtpUsername string `env:"SMTP_USERNAME"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}

type RabbitMQ struct {
	ServerURL  string `env:"RABBITMQ_SERVER_URL"`
	MailsQueue string `env:"RABBITMQ_MAILS_QUEUE"`
}
