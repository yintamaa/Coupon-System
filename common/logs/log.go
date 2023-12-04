package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(filePath string) error {
	logrus.SetLevel(logrus.InfoLevel) // 在测试环境中设置低等级级别，
	writer, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logrus.SetOutput(writer)
	return nil
}
