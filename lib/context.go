package lib

import (
	"os"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

type contextKey int

const (
	uuidKey contextKey = iota
	processorKey
	componentKey
)

func contextFromUUID(ctx context.Context, uuid string) context.Context {
	return context.WithValue(ctx, uuidKey, uuid)
}

func contextFromProcessor(ctx context.Context, processor string) context.Context {
	return context.WithValue(ctx, processorKey, processor)
}

func contextFromComponent(ctx context.Context, component string) context.Context {
	return context.WithValue(ctx, componentKey, component)
}

func uuidFromContext(ctx context.Context) (string, bool) {
	uuid, ok := ctx.Value(uuidKey).(string)
	return uuid, ok
}

func processorFromContext(ctx context.Context) (string, bool) {
	processor, ok := ctx.Value(processorKey).(string)
	return processor, ok
}

func componentFromContext(ctx context.Context) (string, bool) {
	component, ok := ctx.Value(componentKey).(string)
	return component, ok
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	entry := logrus.NewEntry(logrus.New()).WithField("pid", os.Getpid())

	if uuid, ok := uuidFromContext(ctx); ok {
		entry = entry.WithField("uuid", uuid)
	}

	if processor, ok := processorFromContext(ctx); ok {
		entry = entry.WithField("processor", processor)
	}

	if component, ok := componentFromContext(ctx); ok {
		entry = entry.WithField("component", component)
	}

	return entry
}