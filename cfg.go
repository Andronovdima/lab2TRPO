
package apiserver

type Config struct {
	BindAddr    string
	LogLevel    string
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8000",
		LogLevel:    "rc",
		DatabaseURL: "host=localhost dbname=avito-chat sslmode=disable port=5431 password=1234 user=d",
	}
}
