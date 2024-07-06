# golang-log

## log with library zap
https://github.com/uber-go/zap

## log with library logrus
https://github.com/sirupsen/logrus

## log with library gozero
https://github.com/zeromicro/go-zero

### một vài điểm khác biệt
- level: `gozero` chỉ có 4 level: `debug`, `info`, `error`, `Severe`
  - khác với `zap`, `logrus`, thì không có level `WarnLevel`
- Với 1 level thì có 1 func:
  - Info: input là string
  - Infof: input là string và %s
  - Infov: input có thể là 1 struct, kiểu json
  - Infow: input là string và log field
- AddGlobalFields: cực kỳ phù hợp với những api, chức năng có độ phức tạp cao, sẽ note chi tiết trong phần sau.

#### Ví dụ:
1. Code:
```go
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
```
2. Log
```log
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:8","content":"go-zero-log: level Info","level":"info"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:9","content":"go-zero-log: level Infof string data","level":"info"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:15","content":{"name":"111","value":"222"},"level":"info"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:17","content":"go-zero-log: level infow","key":"abc","level":"info"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:23","content":"go-zero-log: level Debug","key":"abc","level":"debug"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:24","content":"go-zero-log: level Debugf string data","key":"abc","level":"debug"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:26","content":"go-zero-log: level Error","key":"abc","level":"error"}
{"@timestamp":"2024-07-06T18:07:42.649+07:00","caller":"gozero/zero.go:27","content":"go-zero-log: level Errorf string data","key":"abc","level":"error"}
```

## example
- run `go run main.go`
```go
package main

import (
	"fmt"

	loggerRus "github.com/ducnpdev/golang-log/logrus"
	loggerZap "github.com/ducnpdev/golang-log/zap"
)

func init() {

}

func main() {
	fmt.Println("main")
	lrus := loggerRus.New()
	lrus.Debugf("xin chao logrus")
}
```
- output:
    ```log
    {"level":"debug","msg":"logrus: log debug","time":"2024-06-16 10:07:45"}
    {"level":"error","msg":"logrus: log error","time":"2024-06-16 10:07:45"}
    {"level":"info","msg":"logrus: log info","time":"2024-06-16 10:07:45"}
    {"level":"warning","msg":"logrus: log warn","time":"2024-06-16 10:07:45"}

    {"LEVEL":"debug","TIME":"2024-06-16T10:07:45.818+0700","CALLER":"golang-log/main.go:23","MESSAGE":"zap: log debug"}
    {"LEVEL":"error","TIME":"2024-06-16T10:07:45.818+0700","CALLER":"golang-log/main.go:24","MESSAGE":"zap: log error"}
    {"LEVEL":"info","TIME":"2024-06-16T10:07:45.818+0700","CALLER":"golang-log/main.go:25","MESSAGE":"zap: log info"}
    {"LEVEL":"warn","TIME":"2024-06-16T10:07:45.818+0700","CALLER":"golang-log/main.go:26","MESSAGE":"zap: log warn"}
    ```

## add log with context
### integration with logrus
