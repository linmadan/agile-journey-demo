package constant

import (
	"github.com/linmadan/agile-journey-demo/pkg/constant"
	"github.com/linmadan/egglib-go/log"
	"github.com/linmadan/egglib-go/log/logrus"
)

var Logger log.Logger

func init() {
	Logger = logrus.NewLogrusLogger()
	Logger.SetServiceName(constant.SERVICE_NAME)
	Logger.SetLevel(constant.LOG_LEVEL)
}
