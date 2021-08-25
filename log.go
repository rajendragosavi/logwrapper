package logwrapper

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	RequestID int
	Message   string
	Status    string
	Error     error
}

// we will intialize the logger object here with all required details and use this as a wrapper package.
func NewLogger() zerolog.Logger {
	log.Logger = log.With().Caller().Logger()                      //this will add file and line number to logs.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //To log a human-friendly, colorized output,

	return log.Logger
}
