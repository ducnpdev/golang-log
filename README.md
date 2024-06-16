# golang-log

## log with library zap
https://github.com/uber-go/zap

## log with library logrus
https://github.com/sirupsen/logrus

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

	lzap := loggerZap.NewLogger("", "", "", "")
	lzap.Debugf("xin chao")
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