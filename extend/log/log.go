package log

import (
	"config-server-go/common"
	"os"

	"github.com/sirupsen/logrus"
)

var fp = common.FlagParams

func init() {
	setLevel()
	setFormatter()
	setLogFile()

	if fp.LogShowLineNumber {
		logrus.AddHook(new(LineHook))
	}

}

func setFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func setLevel() {
	level, _ := logrus.ParseLevel(fp.LogLevel)
	logrus.SetLevel(level)
}

func setLogFile() {
	if fp.LogFile != "" {
		file, err := os.OpenFile(fp.LogFile, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			logrus.Info("Failed to log to file, using default stderr")
		}
	}
}
