package logger

import (
	"bytes"
	"context"
	"io"
	"testing"
)

// Tests that's data consistent for each level
func TestEntry_Log(t *testing.T) {
	out := new(bytes.Buffer)
	meta := meta{
		RequestID: "request-id-1000",
		TraceID:   "trace-id-1000",
		UserID:    "user-id-1000",
	}

	e1 := e(DebugLevel, out, meta)

	c2 := Fields{
		"entry-2": "value-2",
		"entry-2-map": map[string]interface{}{
			"123": "123",
		},
	}

	for _, level := range levels {
		e2 := e1.WithFields(c2)

		// Must be logged content from e2
		e2.Error("e2 message")
		if record, err := RecordFromJSON(out.Bytes()); err == nil {
			record.assert(t, c2)
			record.Level = ErrorLevel.String()
		} else {
			t.Fatal(err)
		}
		out.Reset()

		// Must be logged empty content
		e1.log(level, "e1 message")

		if record, err := RecordFromJSON(out.Bytes()); err == nil {
			record.assert(t, Fields{})
		} else {
			t.Fatal(err)
		}

		out.Reset()
	}
}

func TestEntry_Verbosity_HaveToBeLogged(t *testing.T) {
	out := new(bytes.Buffer)

	// key - verbosity level, value - which levels should be written to the log
	expected := make(map[Level][]Level)

	for i := 0; i < len(levels); i++ {
		level := levels[i]
		expected[level] = levels[:i+1]
	}

	for verbosity, levels := range expected {
		entry := NewLoggerFactory(verbosity, SetOut(out)).
			Logger(context.TODO()).(*entry)

		var missed []Level

		for _, level := range levels {
			entry.log(level, "Success")

			if out.Cap() == 0 {
				missed = append(missed, level)
			}

			out.Reset()
		}

		if len(missed) > 0 {
			t.Errorf("Verbosity '%s' levels must be presented: %+v", verbosity, missed)
		}
	}
}

func TestEntry_Verbosity_HaveToBeIgnored(t *testing.T) {
	out := new(bytes.Buffer)

	// key - verbosity level, value - which levels should be ignored
	unexpected := make(map[Level][]Level)

	for i := 0; i < len(levels); i++ {
		level := levels[i]
		unexpected[level] = levels[i+1:]
	}

	for verbosity, levels := range unexpected {
		entry := NewLoggerFactory(verbosity, SetOut(out)).
			Logger(context.TODO()).(*entry)

		var unexpected []Level

		for _, level := range levels {
			entry.log(level, "Success")

			if out.Cap() != 0 {
				unexpected = append(unexpected, level)
			}

			out.Reset()
		}

		if len(unexpected) > 0 {
			t.Errorf("Verbosity '%s', levels must be ignored: %+v", verbosity, unexpected)
		}
	}
}

func TestSetDefaultEntry(t *testing.T) {
	buf := &bytes.Buffer{}
	entry := NewLoggerFactory(DebugLevel, SetOut(buf)).Logger(context.Background())
	SetDefaultEntry(entry)
	Debug("Hello World!")
	if buf.Len() == 0 {
		t.Fatal("Buffer is empty")
	}

	buf.Reset()
	Debugf("test %s", "test")
	if buf.Len() == 0 {
		t.Fatal("Buffer is empty")
	}
}

func TestDefaultEntry_Immutable(t *testing.T) {
	e := WithFields(Fields{
		"TEST": "TEST",
	}).(*entry)

	if len(e.content) == 0 {
		t.Error("Content is empty")
	}

	if len(logger.(*entry).content) > 0 {
		t.Error("Default entry content has been changed")
	}
}

func e(level Level, out io.Writer, meta meta) *entry {
	factory := NewLoggerFactory(level, SetOut(out))
	entry := factory.Logger(context.TODO()).(*entry)
	entry.meta = meta

	return entry
}
