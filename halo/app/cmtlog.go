package app

import (
	"context"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtlog "github.com/cometbft/cometbft/libs/log"
)

var _ cmtlog.Logger = (*cmtLogger)(nil)

const (
	levelError = iota + 1
	levelInfo
	levelDebug
)

// levels maps strings to numbers for easy comparison.
//
//nolint:gochecknoglobals // Global is ok here.
var levels = map[string]int{
	"error": levelError,
	"info":  levelInfo,
	"debug": levelDebug,
}

// cmtLogger implements cmtlog.Logger by using the omni logging pattern.
// Comet log level is controlled separately in config.toml, since comet logs are very noisy.
type cmtLogger struct {
	ctx   context.Context //nolint:containedctx // This is a wrapper around the omni logger which is context based.
	level int
}

func newCmtLogger(ctx context.Context, levelStr string) (cmtLogger, error) {
	level, ok := levels[strings.ToLower(levelStr)]
	if !ok {
		return cmtLogger{}, errors.New("invalid comet log level", "level", levelStr)
	}

	return cmtLogger{
		ctx:   ctx,
		level: level,
	}, nil
}

func (c cmtLogger) Debug(msg string, keyvals ...any) {
	if c.level < levelDebug {
		return
	}
	log.Debug(c.ctx, msg, keyvals...)
}

func (c cmtLogger) Info(msg string, keyvals ...any) {
	if c.level < levelInfo {
		return
	}
	log.Info(c.ctx, msg, keyvals...)
}

func (c cmtLogger) Error(msg string, keyvals ...any) {
	if c.level < levelError {
		return
	}

	keyvals, err := splitOutError(keyvals)

	log.Error(c.ctx, msg, err, keyvals...)
}

func (c cmtLogger) With(keyvals ...any) cmtlog.Logger { //nolint:ireturn // This signature is required by interface.
	return cmtLogger{
		ctx:   log.WithCtx(c.ctx, keyvals...),
		level: c.level,
	}
}

// splitOutError splits the keyvals into a slice of keyvals without the error and the error.
func splitOutError(keyvals []any) ([]any, error) {
	var remaining []any
	var err error
	for i := 0; i < len(keyvals); i += 2 {
		if keyErr, ok := keyvals[i+1].(error); ok {
			err = keyErr
		} else {
			remaining = append(remaining, keyvals[i], keyvals[i+1])
		}
	}

	return remaining, err
}
