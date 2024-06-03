package injection

import (
	"dudos/arguments"
	"dudos/config"
	payload2 "dudos/payload"
)

func CollectDependency(cfg *config.Config, flag *arguments.Flag) {
	payload, err := payload2.New(flag)
	if err != nil {
		panic(err)
	}

}
