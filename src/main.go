package main

import (
	_ "embed"

	"github.com/triole/logseal"
)

//go:embed conf.toml
var configString string

func main() {
	parseArgs()
	lg = logseal.Init(
		CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON,
	)
	readConfig()
	process()
}
