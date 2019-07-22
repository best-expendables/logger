package logmock

import "bitbucket.org/snapmartinc/logger"

var _ = &Entry{}

type Entry struct {
	DebugFn     func(args ...interface{})
	AlertFn     func(args ...interface{})
	InfoFn      func(args ...interface{})
	NoticeFn    func(args ...interface{})
	WarningFn   func(args ...interface{})
	ErrorFn     func(args ...interface{})
	EmergencyFn func(args ...interface{})
	CriticalFn  func(args ...interface{})

	DebugfFn     func(format string, args ...interface{})
	AlertfFn     func(format string, args ...interface{})
	InfofFn      func(format string, args ...interface{})
	NoticefFn    func(format string, args ...interface{})
	WarningfFn   func(format string, args ...interface{})
	ErrorfFn     func(format string, args ...interface{})
	EmergencyfFn func(format string, args ...interface{})
	CriticalfFn  func(format string, args ...interface{})

	WithFieldFn  func(key string, value interface{}) logger.Entry
	WithFieldsFn func(fields logger.Fields) logger.Entry
}

func (e *Entry) Debug(args ...interface{}) {
	e.DebugFn(args)
}

func (e *Entry) Alert(args ...interface{}) {
	e.AlertFn(args)
}

func (e *Entry) Info(args ...interface{}) {
	e.InfoFn(args)
}

func (e *Entry) Notice(args ...interface{}) {
	e.NoticeFn(args)
}

func (e *Entry) Warning(args ...interface{}) {
	e.WarningFn(args)
}

func (e *Entry) Error(args ...interface{}) {
	e.ErrorFn(args)
}

func (e *Entry) Emergency(args ...interface{}) {
	e.EmergencyFn(args)
}

func (e *Entry) Critical(args ...interface{}) {
	e.CriticalFn(args)
}

func (e *Entry) WithField(key string, value interface{}) logger.Entry {
	return e.WithFieldFn(key, value)
}

func (e *Entry) WithFields(fields logger.Fields) logger.Entry {
	return e.WithFieldsFn(fields)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.DebugfFn(format, args)
}

func (e *Entry) Alertf(format string, args ...interface{}) {
	e.AlertfFn(format, args)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.InfofFn(format, args)
}

func (e *Entry) Noticef(format string, args ...interface{}) {
	e.NoticefFn(format, args)
}

func (e *Entry) Warningf(format string, args ...interface{}) {
	e.WarningfFn(format, args)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.ErrorfFn(format, args)
}

func (e *Entry) Emergencyf(format string, args ...interface{}) {
	e.EmergencyfFn(format, args)
}

func (e *Entry) Criticalf(format string, args ...interface{}) {
	e.CriticalfFn(format, args)
}
