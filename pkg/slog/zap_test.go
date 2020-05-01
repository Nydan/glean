package slog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestGetZapLevel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(zapcore.InfoLevel, getZapLevel("info"))
	assert.Equal(zapcore.WarnLevel, getZapLevel("warn"))
	assert.Equal(zapcore.DebugLevel, getZapLevel("debug"))
	assert.Equal(zapcore.ErrorLevel, getZapLevel("error"))
	assert.Equal(zapcore.FatalLevel, getZapLevel("fatal"))
}

func TestGetZapEncoding(t *testing.T) {
	_ = getEncoder(false)
}

func TestNewZapWithConfigFile(t *testing.T) {
	_ = newZapLogger(Configuration{EnableFile: true})
}

func TestZapFatalw(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Fatal not called")
		}
	}()

	zap := zapLogger{nil}
	zap.Fatalw("something")

}
