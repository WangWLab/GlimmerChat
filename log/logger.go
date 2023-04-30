package log

import (
	"os"

	"github.com/op/go-logging"
)

var RootLogger = logging.MustGetLogger("root")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{callpath} ▶ %{level:.4s} %{color:reset} %{message}`,
)

func init() {
	rootBackend := logging.NewLogBackend(os.Stdout, "", 0)
	rootBackendFormatter := logging.NewBackendFormatter(rootBackend, format)
	rootBackendLeveled := logging.AddModuleLevel(rootBackendFormatter)
	rootBackendLeveled.SetLevel(logging.DEBUG, "")
	RootLogger.SetBackend(rootBackendLeveled)
	RootLogger.Warning("Logger initialized")
}
