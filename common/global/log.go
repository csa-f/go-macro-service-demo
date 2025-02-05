package global

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/csa-f/macro-service/common/config"
	log "github.com/sirupsen/logrus"
)

func InitLog(conf *config.Log) {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&LogFormat{
		Color: &conf.Color,
	})
	log.SetLevel(conf.Level)
}

type LogFormat struct {
	Color *config.LogColor
}

func (f *LogFormat) Format(entry *log.Entry) ([]byte, error) {
	var levelColor string
	switch entry.Level {
	case log.DebugLevel:
		levelColor = f.Color.Debug
	case log.TraceLevel:
		levelColor = f.Color.Trace
	case log.WarnLevel:
		levelColor = f.Color.Warn
	case log.ErrorLevel:
		levelColor = f.Color.Error
	case log.FatalLevel:
		levelColor = f.Color.Fatal
	case log.PanicLevel:
		levelColor = f.Color.Panic
	case log.InfoLevel:
		levelColor = f.Color.Info
	}

	buffer := bytes.Buffer{}
	buffer.WriteString(time.DateTime)
	buffer.WriteString(",000")

	funcs := strings.SplitN(entry.Caller.Function, ".", 2)
	msg := fmt.Sprintf(
		"\033[%sm[%s] [%s] [%s:%d] %s\033[0m \n",
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Time.Format(buffer.String()),
		funcs[0]+"/"+filepath.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message,
	)
	return []byte(msg), nil
}
