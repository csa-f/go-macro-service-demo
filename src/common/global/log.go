package global

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/csa-f/go-macro-service-demo/common/config"
	log "github.com/sirupsen/logrus"
)

func InitLog(conf *config.Log) {
	level := conf.Level
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&LogFormat{
		red:    conf.Color.Red,
		yellow: conf.Color.Yellow,
		gray:   conf.Color.Def,
		def:    conf.Color.Def,
		green:  conf.Color.Green,
	})
	log.SetLevel(log.Level(level))
}

type LogFormat struct {
	red    string
	yellow string
	gray   string
	def    string
	green  string
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
	case log.DebugLevel,
		log.TraceLevel:
		levelColor = green
	case log.WarnLevel:
		levelColor = yellow
	case log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel:
		levelColor = red
	case log.InfoLevel:
		levelColor = def
	default:
		levelColor = def
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
