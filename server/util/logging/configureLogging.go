package logging

import (
	"os"
	"time"

	"ordo-map/settings"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func Configure() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.ErrorFieldName = "e"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldUnit = time.Duration(zerolog.DurationFieldUnit.Milliseconds())
	zerolog.DurationFieldInteger = true
	if settings.Current.Logging.LogLevel == "Debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if settings.Current.Logging.LogLevel == "Info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if settings.Current.Logging.LogLevel == "Warn" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if settings.Current.Logging.LogLevel == "Error" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if settings.Current.Logging.LogLevel == "Fatal" {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	} else if settings.Current.Logging.LogLevel == "Panic" {
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Logger = log.With().Caller().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Logging initialized")
}
