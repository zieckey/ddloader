package ddl

import (
    "github.com/zieckey/goini"
)

type AlarmLevel int

const (
    AlarmLevelInfo AlarmLevel = iota
    AlarmLevelWarn
    AlarmLevelError
    AlarmLevelFatal
)

// 参数ini是配置文件解析出来的数据结构，上层应用可以在这里定义更多跟报警相关配置项，例如：
//  a. 报警服务地址
//  b. 报警项目id
// 等等与报警平台相关的特殊配置项
type AlarmCallback func(msg string, l AlarmLevel, ini* goini.INI) error

func DefaultAlarmCallback(msg string, l AlarmLevel, ini* goini.INI) error {
    switch l {
    case AlarmLevelInfo:
        Infof("%s", msg)
        return nil
    case AlarmLevelWarn:
        Warnf("%s", msg)
        return nil
    case AlarmLevelError:
        Errorf("%s", msg)
        return nil
    case AlarmLevelFatal:
        Fatalf("%s", msg)
        return nil
    }
    return nil
}

var defaultAlarmCallback AlarmCallback

func SetAlarm(f AlarmCallback) {
    defaultAlarmCallback = f
}