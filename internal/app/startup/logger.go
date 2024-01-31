package startup

import (
	logger "gateway/pkg"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	return logger.NewLogger()
}
