package debug

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func LogError(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Error().Msgf("Error: %s occurred at %s:%d", err.Error(), file, line)
}