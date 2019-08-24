package logger

import "github.com/alexcesaro/log/stdlog"

var Logger = stdlog.GetFromFlags()

// facade for stdlog
var Info = Logger.Info
var Warning = Logger.Warning
var Debug = Logger.Debug
var Error = Logger.Error
