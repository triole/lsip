package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/triole/logseal"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "lsip"
	appDescription = "list your ip"
	appMainversion = "0.1"
	lg             = logseal.Init("info", nil, false, false)
)

var CLI struct {
	Print       string `help:"print certain ip version, can be: 4, 6, both, any or all" short:"p" enum:"4,6,both,any,all" default:"any"`
	Threads     int    `help:"threads, max of parallel requests" short:"t" default:"16"`
	DryRun      bool   `help:"dry run, just print don't do" short:"n"`
	LogFile     string `help:"log file" default:"/dev/stdout"`
	LogLevel    string `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text"`
	LogJSON     bool   `help:"enable json log, instead of text one"`
	VersionFlag bool   `help:"display version" short:"V"`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}
