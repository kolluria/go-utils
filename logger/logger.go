package logger

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New(ctx context.Context) zerolog.Logger {
	zerolog.Ctx(ctx)
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}

func NewWithContext(ctx context.Context, w *io.Writer) Logger {
	var logger zerolog.Logger
	if w != nil {
		logger = zerolog.New(*w)
	} else {
		logger = zerolog.New(os.Stdout)
	}
	// Attach the Logger to the context.Context
	ctx = logger.WithContext(ctx)
	someFunc(ctx)
}
