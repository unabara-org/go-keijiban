package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/unabara-org/go-keijiban/utils"
	"os"
)

func NewFileLogger() (echo.Logger, func() error) {
	path := "storage/logs/log"

	var file *os.File
	var err error

	if utils.ExistsFile(path) {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			panic("ログファイルのオープンに失敗しました")
		}
	} else {
		file, err = os.Create(path)
		if err != nil {
			panic("ログファイルの生成に失敗しました")
		}
	}

	l := log.New("")
	l.SetOutput(file)

	return l, func() error {
		return file.Close()
	}
}
