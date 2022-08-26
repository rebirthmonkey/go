package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/marmotedu/iam/pkg/log/klog"
)

type Config struct {
	OutputPaths       []string
	ErrorOutputPaths  []string
	Level             string
	Format            string
	DisableCaller     bool
	DisableStacktrace bool
	EnableColor       bool
	Development       bool
	Name              string
}

type CompletedConfig struct {
	*Config
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

func (c CompletedConfig) New() error {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(c.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	// when output to local path, with color is forbidden
	if c.Format == consoleFormat && c.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       c.Development,
		DisableCaller:     c.DisableCaller,
		DisableStacktrace: c.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         c.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      c.OutputPaths,
		ErrorOutputPaths: c.ErrorOutputPaths,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := &zapLogger{
		zapLogger: l.Named(c.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}

	klog.InitLogger(l)
	zap.RedirectStdLog(l)

	std = logger

	return nil
}
