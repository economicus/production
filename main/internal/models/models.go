package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var logger log.Logger

type format struct {
	log.TextFormatter
}

func (f *format) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] - %s - %s\n", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&format{
		log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
	})
}
