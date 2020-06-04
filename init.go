package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	env     = environment{}
	setting = config{}
)

func init() {
	/* Set Input ENV */
	env.configpath = os.Getenv("CONFIG_LOCATION")
	if len(env.configpath) == 0 {
		env.configpath = "./etc/go-server/config.d/setting.conf"
	}
	env.logdir = os.Getenv("LOG_DIR_LOCATION")
	if len(env.logdir) == 0 {
		env.logdir = "./var/log/go-server/"
	}
	/* Set Input Config */
	body, err := ioutil.ReadFile(env.configpath)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := yaml.Unmarshal(body, &setting); err != nil {
		logrus.Fatal(err)
	}

	/* create a formatter */
	formatter := new(logrus.TextFormatter)
	formatter.DisableColors = setting.Logger.DisableColors
	formatter.ForceQuote = setting.Logger.ForceQuote
	formatter.FullTimestamp = setting.Logger.FullTimeStamp
	formatter.PadLevelText = setting.Logger.PadLevelText
	formatter.QuoteEmptyFields = setting.Logger.QuoteEmptyFields
	formatter.TimestampFormat = setting.Logger.TimeStampFormat /*"2006-01-02 15:04:05"*/

	/* Set Logger Config */
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.TraceLevel)

	if len(env.logdir) == 0 {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.AddHook(&WriterHook{
			Writer: newLoggerRotate(fmt.Sprintf("%v/errors.log", env.logdir)),
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
			},
		})
		logrus.AddHook(&WriterHook{
			Writer: newLoggerRotate(fmt.Sprintf("%v/processing_tasks.log", env.logdir)),
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
			},
		})
		logrus.AddHook(&WriterHook{
			Writer: newLoggerRotate(fmt.Sprintf("%v/warnings.log", env.logdir)),
			LogLevels: []logrus.Level{
				logrus.WarnLevel,
			},
		})
		logrus.AddHook(&WriterHook{
			Writer: newLoggerRotate(fmt.Sprintf("%v/debugs.log", env.logdir)),
			LogLevels: []logrus.Level{
				logrus.DebugLevel,
			},
		})
	}
}
