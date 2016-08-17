package cmd

import (
	"errors"
	"flag"
	"path/filepath"
	"syscall"

	"github.com/cybozu-go/log"
)

var (
	logFilename = flag.String("logfile", "", "Log filename")
	logLevel    = flag.String("loglevel", "info", "Log level [critical,error,warning,info,debug]")
	logFormat   = flag.String("logformat", "plain", "Log format [plain,logfmt,json]")
)

// LogConfig configures cybozu-go/log's default logger.
//
// Filename, if not an empty string, specifies the output filename.
//
// Level is the log threshold level name.
// Valid levels are "critical", "error", "warning", "info", and "debug".
// Empty string is treated as "info".
//
// Format specifies log formatter to be used.
// Available formatters are "plain", "logfmt", and "json".
// Empty string is treated as "plain".
//
// For details, see https://godoc.org/github.com/cybozu-go/log .
type LogConfig struct {
	Filename string `toml:"filename" json:"filename"`
	Level    string `toml:"level"    json:"level"`
	Format   string `toml:"format"   json:"format"`
}

// Apply applies configurations to the default logger.
//
// Command-line flags take precedence over the struct member values.
func (c *LogConfig) Apply() error {
	logger := log.DefaultLogger()

	filename := c.Filename
	if flag.Lookup("logfile") != nil {
		filename = *logFilename
	}
	if len(filename) > 0 {
		abspath, err := filepath.Abs(filename)
		if err != nil {
			return err
		}
		w, err := log.NewFileReopener(abspath, syscall.SIGUSR1)
		if err != nil {
			return err
		}
		logger.SetOutput(w)
	}

	level := c.Level
	if flag.Lookup("loglevel") != nil {
		level = *logLevel
	}
	if len(level) == 0 {
		level = "info"
	}
	err := logger.SetThresholdByName(level)
	if err != nil {
		return err
	}

	format := c.Format
	if flag.Lookup("logformat") != nil {
		format = *logFormat
	}
	switch format {
	case "", "plain":
		logger.SetFormatter(log.PlainFormat{})
	case "logfmt":
		logger.SetFormatter(log.Logfmt{})
	case "json":
		logger.SetFormatter(log.JSONFormat{})
	default:
		return errors.New("invalid format: " + format)
	}

	return nil
}
