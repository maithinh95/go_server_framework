package main

import (
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func newLoggerRotate(filename string) *lumberjack.Logger {
	logger := &lumberjack.Logger{
		Filename: filename,
		MaxSize:  50, // megabytes
		// MaxBackups: 3,
		// MaxAge:     29, //days
		LocalTime: true,
		Compress:  true, // disabled by default
	}
	go func() {
		for {
			delay := func() time.Duration {
				now := time.Now()
				begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
				return 24*time.Hour - now.Sub(begin)
			}()
			logrus.Debugf("LogRotate for %v delay in %v ", filename, delay)
			timer := time.NewTimer(delay)
			select {
			case t := <-timer.C:
				if err := logger.Rotate(); err != nil {
					logrus.Error(err)
				}
				logrus.Debugf("New logger at %v", t.String())
			}
		}
	}()
	return logger
}
