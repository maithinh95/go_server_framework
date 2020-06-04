package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	/* TEST */
	logrus.Infof("Environment= %+v", env)
	logrus.Infof("Setting= %+v", setting)

	for {
		now := time.Now().Local().String()
		logrus.Infof(" NOW= %v", now[:21])
		logrus.Error(" ERROR_TEST")
		time.Sleep(3 * time.Second)
	}
}
