package logmock

import "context"

var _ = Factory{}

type Factory struct {
	LoggerFn func(ctx context.Context) Entry
}

func (p Factory) Logger(ctx context.Context) Entry {
	return p.LoggerFn(ctx)
}
