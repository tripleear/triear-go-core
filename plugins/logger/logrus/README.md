# logrus

[logrus](https://github.com/sirupsen/logrus) logger implementation for __triear-go__ [meta logger](https://github.com/tripleear/triear-go-core/tree/master/logger).

## Usage

```go
import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/tripleear/triear-go-core/logger"
)

func ExampleWithOutput() {
	logger.DefaultLogger = NewLogger(logger.WithOutput(os.Stdout))
	logger.Infof("testing: %s", "Infof")
}

func ExampleWithLogger() {
	l := logrus.New() // *logrus.Logger
	logger.DefaultLogger = NewLogger(WithLogger(l))
	logger.Infof("testing: %s", "Infof")
}
```

