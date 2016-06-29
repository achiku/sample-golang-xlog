package main

import (
	"log"

	"github.com/rs/xlog"
)

func main() {
	conf := xlog.Config{
		Level:  xlog.LevelInfo,
		Output: xlog.NewConsoleOutput(),
		Fields: xlog.F{"init": "mymodule"},
	}
	logger := xlog.New(conf)
	xlog.SetLogger(logger)

	xlog.Info("hello", xlog.F{"func": "main"})
	xlog.Info("bye")

	// want to set tmp logger
	conf.Fields = xlog.F{"extra": "hello"}
	log.Println(conf.Fields)
	xlog.Info("hey")

	// using logger
	logger.Info("from logger 1")
	logger.SetField("key", "oo")
	logger.Info("from logger 2")
	logger.Info("from logger 3")
	xlog.Info("from xlog.Info")

	l := xlog.Copy(logger)
	l.SetField("device", "apple")
	l.Info("aaa")
	logger.Info("aaa")
}
