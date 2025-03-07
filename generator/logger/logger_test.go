package logger_test

import (
	"testing"

	"github.com/djpiper28/setmoji/generator/logger"
	"github.com/stretchr/testify/assert"
)

func TestLoggerIsNotNil(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, logger.Logger)
	logger.Logger.Info("testing 123")
}
