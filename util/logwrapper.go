package Util

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var ZLog *zap.Logger

func init() {
	logDir := "./log/"
	logLevel := "debug"

	_ = logLevel

	srvName := logDir + filepath.Base(os.Args[0])
	pos := strings.Index(srvName, ".exe")
	if pos != -1 {
		srvName = srvName[:pos]
	}
	//	srvName := logdir + srvType + "." + srvID
	ZLog = NewLogger(srvName+".log", true)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
func UniqsTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, "20060102 15:04:05.000", enc)
}

func NewLogger(logFileName string, logToConsole bool) *zap.Logger {
	cfg := zap.NewProductionConfig()

	// set log level
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.LevelKey = "LV"
	cfg.EncoderConfig.TimeKey = "MS"
	cfg.EncoderConfig.CallerKey = "BT"
	cfg.EncoderConfig.MessageKey = "MSG"
	//cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	//{"level":"debug","ts":"2020-09-26T15:21:35.1503541+08:00"
	//cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	//{"level":"debug","ts":"2020-09-26T15:19:02+08:00"
	//cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//{"level":"debug","ts":"2020-09-26T15:19:52.403+0800"
	//cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	//{"level":"debug","ts":1601104814871.312
	//cfg.EncoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	//{"level":"debug","ts":1601104841219876000
	//cfg.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	//{"level":"debug","ts":1601104866.6949341
	cfg.EncoderConfig.EncodeTime = UniqsTimeEncoder

	if logFileName != "" {
		// create directory if needed.
		pos := strings.LastIndex(logFileName, "/")
		if pos != -1 {
			logFolder := logFileName[0:pos]
			var logFilePerm os.FileMode = 0644

			if logFolder == "" {
				logFolder = "./"
			} else {
				if !strings.HasSuffix(logFolder, "/") {
					logFolder += "/"
				}
				err := os.MkdirAll(logFolder, logFilePerm)
				if err != nil {
					logFileName = "./" + logFileName[pos:]
					fmt.Println("log_wrapper InitLog os.MkdirAll failed. err:", err, " using current dir instead. logFileName:", logFileName)
				}
			}
		}
		cfg.OutputPaths = []string{
			logFileName,
		}

		if logToConsole {
			cfg.OutputPaths = append(cfg.OutputPaths, "stdout")
		}
	} else {
		// force log to console
		if !logToConsole {
			fmt.Println("Warn: zap log NewLogger. as logFileName is empty, logToConsole is set to true")
		}

		cfg.OutputPaths = []string{
			"stdout",
		}
	}

	log, err := cfg.Build()
	if err != nil {
		fmt.Println("NewLogger cfg.Build zap log got some error. err:", err)
	}

	defer log.Sync()
	return log
}
