package slogpretty

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	stdLog "log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func SetupPrettyLogger() *slog.Logger {

	loggerOptions := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	return slog.New(loggerOptions.NewPrettyHandler(os.Stdout))
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

type Attr struct {
	key   string
	value interface{}
}

const (
	Enter byte = 10
	Space byte = 32
)

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make([]Attr, 0, r.NumAttrs()+len(h.attrs))

	for _, a := range h.attrs {
		fields = append(fields, Attr{key: a.Key, value: a.Value.Any()})
	}

	r.Attrs(func(a slog.Attr) bool {
		fields = append(fields, Attr{key: a.Key, value: a.Value.Any()})
		return true
	})

	var Attributes []byte

	for _, field := range fields {
		FormatField := fmt.Sprintf("%s: %s"+",", field.key, field.value)
		b, err := json.Marshal(FormatField)
		if err != nil {
			return err
		}
		Attributes = append(Attributes, Enter, Space, Space)
		Attributes = append(Attributes, b...)
	}

	timeStr := r.Time.Format("[15:04:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(Attributes)),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
