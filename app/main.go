package main

import (
	"dudos/arguments"
	"dudos/config"
	"dudos/sender"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetFormatter(
		&logrus.JSONFormatter{
			DisableTimestamp:  false,
			DisableHTMLEscape: true,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", frame.File, frame.Line)
			},
			PrettyPrint: true,
		})
	logrus.SetReportCaller(true)
	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}

	flag1 := arguments.New()

	sender.New()

	logrus.SetLevel(lvl)
	logrus.Fatalf("%+v, %+v", cfg, flag1)
}
