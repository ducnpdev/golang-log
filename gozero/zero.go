package zlogger

import (
	logZero "github.com/zeromicro/go-zero/core/logx"
)

func LoggerZero() {
	logZero.Info("go-zero-log: level Info")
	logZero.Infof("go-zero-log: level Infof %s", "string data")
	type JsonValue struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	var jsonExam = JsonValue{Name: "111", Value: "222"}
	logZero.Infov(jsonExam)

	logZero.Infow("go-zero-log: level infow", logZero.LogField{
		Key: "key", Value: "abc",
	})

	logZero.AddGlobalFields(logZero.LogField{Key: "key", Value: "abc"})

	logZero.Debug("go-zero-log: level Debug")
	logZero.Debugf("go-zero-log: level Debugf %s", "string data")

	logZero.Error("go-zero-log: level Error")
	logZero.Errorf("go-zero-log: level Errorf %s", "string data")

}
