package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log        *zap.Logger
	LOG_LEVEL  = "LOG_LEVEL"
	LOG_OUTPUT = "LOG_OUTPUT"
)

func init() {
	logConfig := zap.Config{ //Cria um log personalizado
		OutputPaths: []string{getOutputLogs()},            // informa o log no terminal
		Level:       zap.NewAtomicLevelAt(getLevelLogs()), //Informa o level do log através de uma variável
		Encoding:    "json",                               //Retorna a informação em json
		EncoderConfig: zapcore.EncoderConfig{ //Configura como vem as informações
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build() // Constrói as informações do logConfig
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutputLogs() string { //Função para chamar aonde vai ser mostrado o log
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if output == "" {
		output = "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level { //Função para escolher qual o level do log
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
