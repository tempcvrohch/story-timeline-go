package global

import "github.com/gin-gonic/gin"

type Config struct {
	Debug    bool
	Server   ServerConfig
	Database DatabaseConfig
	Gemini   GeminiConfig `mapstructure:",squash"`
	Nats     NatsConfig   `mapstructure:",squash"`
}

type NatsConfig struct {
	Host     string `mapstructure:"NATS_HOST"`
	Port     int    `mapstructure:"NATS_PORT"`
	HttpPort int    `mapstructure:"NATS_HTTP_PORT"`
}

type GeminiConfig struct {
	ApiKey      string `mapstructure:"GEMINI_API_KEY"`
	Model       string `mapstructure:"GEMINI_MODEL,omitempty"`
	Temperature float32
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}
type Context struct {
	Config *Config
	Router *gin.Engine
}
