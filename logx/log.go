package logx

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap(level string, production bool) (*zap.Logger, error) {
	var c zap.Config

	if production {
		c = zap.NewProductionConfig()
	} else {
		c = zap.NewDevelopmentConfig()
		c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	if err := c.Level.UnmarshalText([]byte(level)); err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	logger, err := c.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return logger, nil
}
