package logger

import (
	"os"
	"strings"

	"github.com/ekasc/nucleus-api/internal/config"
	"github.com/rs/zerolog"
)

func New(c config.Config) *zerolog.Logger {
	level, err := zerolog.ParseLevel(strings.ToLower(c.LogLevel))
	if err != nil {
		level = zerolog.InfoLevel
	}

	l := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
	return &l
}

func WithReq(l *zerolog.Logger, reqID string) *zerolog.Logger {
	if reqID == "" {
		return l
	}
	logger := l.With().Str("req_id", reqID).Logger()
	return &logger
}
