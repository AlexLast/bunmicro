package bunmicro

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"go-micro.dev/v4/logger"
)

// QueryHook defines the
// structure of our query hook
// it implements the bun.QueryHook
// interface
type QueryHook struct {
	bun.QueryHook

	logger       logger.Logger
	slowDuration time.Duration
}

// QueryHookOptions defines the
// available options for a new
// query hook.
type QueryHookOptions struct {
	Logger       logger.Logger
	SlowDuration time.Duration
}

// NewQueryHook returns a new query hook for use with
// uptrace/bun.
func NewQueryHook(options QueryHookOptions) QueryHook {
	return QueryHook{
		logger:       options.Logger,
		slowDuration: options.SlowDuration,
	}
}

func (qh QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (qh QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	queryDuration := time.Since(event.StartTime)

	// Errors will always be logged
	if event.Err != nil {
		qh.logger.Logf(logger.ErrorLevel, "query: %s, error: %v, duration_ms: %d", event.Query, event.Err, queryDuration.Milliseconds())
		return
	}

	// Queries over a slow time duration will be logged as debug
	if queryDuration >= qh.slowDuration {
		qh.logger.Logf(logger.DebugLevel, "query: %s, duration_ms: %d", event.Query, queryDuration.Milliseconds())
	}
}
