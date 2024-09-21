package global

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/csa-f/go-macro-service-demo/common/conf"
	log "github.com/sirupsen/logrus"
)

func InitLog() {
	level := conf.GetConfig().Server.LogLevel
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(new(LogFormat))
	log.SetLevel(log.Level(level))
}

type LogFormat struct {
}

const (
	red    = "38;2;255;0;0"
	yellow = "38;2;255;165;0"
	gray   = "38;2;100;100;100"
	def    = "38;2;100;150;200"
	green  = "38;2;85;107;47"
)

func (f *LogFormat) Format(entry *log.Entry) ([]byte, error) {
	var levelColor string
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = green
	case log.WarnLevel:
		levelColor = yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = red
	case log.InfoLevel:
		levelColor = def
	default:
		levelColor = def
	}

	funcs := strings.SplitN(entry.Caller.Function, ".", 2)
	msg := fmt.Sprintf(
		"\033[%sm[%s] [%s] [%s:%d] %s\033[0m \n",
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Time.Format("2006-01-02 15:04:05,000"),
		funcs[0]+"/"+filepath.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message,
	)
	return []byte(msg), nil
}
