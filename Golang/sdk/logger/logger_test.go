package logger

import "testing"

func TestSimpleLoggerUse(t *testing.T) {
	SimpleSugarLogger()

	SimpleBasicLogger()

	ConfigureLogger()

	ConfigureJsonLogger()

	MultiLogger()

	LoggerWithColor()
}
