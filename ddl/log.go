package ddl

import (
    "github.com/golang/glog"
)

var defaultLogger Logger

type Logger interface {
    Infof(format string, args ...interface{})
    Warnf(format string, args ...interface{})
    Errorf(format string, args ...interface{})
    Fatalf(format string, args ...interface{})
}

func SetLogger(f Logger) {
    defaultLogger = f
}

type GLog struct {
}

func (g* GLog) Infof(format string, args ...interface{}) {
    glog.Infof(format, args)
}

func (g* GLog) Warnf(format string, args ...interface{}) {
    glog.Warningf(format, args)
}

func (g* GLog) Errorf(format string, args ...interface{}) {
    glog.Errorf(format, args)
}

func (g* GLog) Fatalf(format string, args ...interface{}) {
    glog.Fatalf(format, args)
}

func Infof(format string, args ...interface{}) {
    defaultLogger.Infof(format, args)
}

func Warnf(format string, args ...interface{}) {
    defaultLogger.Warnf(format, args)
}

func Errorf(format string, args ...interface{}) {
    defaultLogger.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
    defaultLogger.Fatalf(format, args)
}