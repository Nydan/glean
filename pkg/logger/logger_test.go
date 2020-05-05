package logger

import (
	"errors"
	"testing"
)

var x Logger

func TestRunTheLog(t *testing.T) {
	NewLogger(Configuration{EnableConsole: true, ConsoleJSONFormat: true})
	x = GetInstance()

	testTable := map[string]func(){
		"Debugf": func() { Debugf("tag") },
		"Infof":  func() { Infof("tag") },
		"Warnf":  func() { Warnf("tag") },
		"Errorf": func() { Errorf("%v", errors.New("boom")) },
		"Infow":  func() { Infow("message", "message", "content message") },
		"Warnw":  func() { Warnw("message", "message", "content message") },
		"Error":  func() { Errorw("message", "message", "content message") },
	}

	for name, tt := range testTable {
		t.Run(name, func(t *testing.T) {
			tt()
		})
	}
}

func TestWithFileds(t *testing.T) {
	WithFields(map[string]interface{}{
		"message": "content message",
	})
}
