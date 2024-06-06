package main

import (
	"dudos/CETb"
	"dudos/arguments"
	"dudos/config"
	payload2 "dudos/payload"
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

	flags := arguments.New()
	if flags == nil {
		logrus.Fatal("flags is nil")
	}

	flags.Payload, err = payload2.New(flags.PayloadType)
	if err != nil {
		logrus.Fatal(err)
	}

	sender.New()

	logrus.SetLevel(lvl)
	cet := CETb.New()
	cet.StartLoad(flags)

	logrus.Fatalf("%+v, %+v", cfg, flags)
}
