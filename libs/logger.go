package libs

import (
	"github.com/labstack/gommon/log"
)

type Logger interface {
	InitLogger()
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

type ApiLogger struct {
}

func NewApiLogger() *ApiLogger {
	return &ApiLogger{}
}

func (l *ApiLogger) InitLogger() {

}

func (l *ApiLogger) Info(message string) {
	log.Info(message)
}

func (l *ApiLogger) Warn(message string) {
	log.Warn(message)
}

func (l *ApiLogger) Error(message string) {
	log.Error(message)
}

func (l *ApiLogger) Fatal(message string) {
	log.Fatal(message)
}
