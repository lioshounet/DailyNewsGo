package bootstarp

import (
	"thor/router"
	"thor/util/helper"
	"thor/util/system"
)

func Run(opts *system.Options) {
	err := helper.LoadConf(opts.ConfFile)
	if err != nil {
		panic(err)
	}

	pid := opts.GetPidFile(helper.GetAppConf().AppName)
	if err := initServer(); err != nil {
		panic(err)
	}

	if err := initGin(pid, router.RegisterRoute); err != nil {
		panic(err)
	}
}
