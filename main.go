package main

import (
	loggerRus "github.com/ducnpdev/golang-log/logrus"
	loggerZap "github.com/ducnpdev/golang-log/zap"
)

func init() {

}

func main() {

	lrus := loggerRus.New()
	lrus.Debugf("logrus: log debug")
	lrus.Errorf("logrus: log error")
	lrus.Infof("logrus: log info")
	lrus.Warnf("logrus: log warn")

	lzap := loggerZap.NewLogger("", "", "", "")
	lzap.Debugf("zap: log debug")
	lzap.Errorf("zap: log error")
	lzap.Infof("zap: log info")
	lzap.Warnf("zap: log warn")
}
