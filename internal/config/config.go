package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           int
	DBDsn          string
	RedisURL       string
	JWTSecret      string
	AllowedOrigins []string
	AppEnv         string
	LogLevel       string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	portStr := strings.TrimSpace(os.Getenv("PORT"))
	port := 8080

	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			return Config{}, fmt.Errorf("invalid port: %w", err)
		}

		port = p

	}

	var allowedOrigins []string

	origins := strings.SplitSeq(strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS")), ",")
	for origin := range origins {

		o := strings.TrimSpace(origin)
		if o != "" {
			allowedOrigins = append(allowedOrigins, o)
		}
	}

	appEnv := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))

	if appEnv == "" {
		appEnv = "dev"
	}

	logLevel := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	dbDsn := strings.TrimSpace(os.Getenv("DB_DSN"))
	redisURL := strings.TrimSpace(os.Getenv("REDIS_URL"))
	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))

	if logLevel == "" {
		logLevel = "info"
	}

	c := Config{
		Port:           port,
		DBDsn:          dbDsn,
		RedisURL:       redisURL,
		JWTSecret:      jwtSecret,
		AllowedOrigins: allowedOrigins,
		AppEnv:         appEnv,
		LogLevel:       logLevel,
	}
	return c, nil
}

func Validate(c Config) error {
	if c.Port <= 0 || c.Port > 65535 {
		return errors.New("env var PORT must be a valid port number (1–65535)")
	}

	if len(c.AllowedOrigins) == 0 && c.AppEnv != "dev" {
		return errors.New("env var ALLOWED_ORIGINS is not set")
	}

	for _, o := range c.AllowedOrigins {
		if o == "*" && c.AppEnv == "prod" {
			return errors.New("wildcard allowed origin is not allowed in production")
		}
	}

	if c.AppEnv != "prod" && c.AppEnv != "dev" && c.AppEnv != "test" {
		return errors.New("env var APP_ENV must be 'dev', 'prod', or 'test'")
	}

	switch c.LogLevel {
	case "debug", "info", "warn", "error":
		// valid
	default:
		return errors.New("env var LOG_LEVEL must be 'debug', 'info', 'warn', or 'error'")
	}

	if c.DBDsn == "" {
		return errors.New("env var DB_DSN is not set")
	}

	if c.JWTSecret == "" || len(c.JWTSecret) < 32 {
		return errors.New("env var JWT_SECRET is not set or too short (min 32 chars)")
	}

	if c.RedisURL == "" {
		return errors.New("env var REDIS_URL is not set")
	}

	return nil
}

func MustLoad() Config {
	c, err := Load()
	if err != nil {
		panic(err)
	}

	err = Validate(c)
	if err != nil {
		panic(err)
	}

	return c
}

func IsProd(c Config) bool {
	return c.AppEnv == "prod"
}
