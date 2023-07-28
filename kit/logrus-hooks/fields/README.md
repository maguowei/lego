# logrus-fields-hook

logrus add `Default Fields` use `Hooks`, without using `log.WithFields` and `logrus.Entry`.

### Usage

```go
package main

import (
	"github.com/maguowei/lego/kit/logrus-hooks/fields"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	appName := "testApp"
	logger.AddHook(fields.NewHook(logrus.Fields{"app": appName}, []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.InfoLevel, logrus.DebugLevel}))
	logger.Info("hello world")
	// output: {"app":"testApp","level":"info","msg":"hello world","time":"2019-01-02T15:04:05+08:00"}
}
```
