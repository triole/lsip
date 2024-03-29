package main

import (
	"github.com/triole/logseal"
)

func process() {
	chin := make(chan tReq, CLI.Threads)
	chout := make(chan tReq, CLI.Threads)

	for _, el := range conf.IPRetrieval {
		go makeReq(el, chin, chout)
	}

	c := 0
	ln := len(conf.IPRetrieval)
	for req := range chout {
		c++
		if CLI.All {
			printLog("got response", req, len(chin))
		} else {
			// default mode to print first fetched ip
			if req.Error == nil && (isValidIPv4(req.IPv4) || isValidIPv6(req.IPv6)) {
				printLog("got response", req, len(chin))
				close(chin)
				close(chout)
				break
			}
		}
		if c >= ln {
			close(chin)
			close(chout)
			break
		}
	}
}

func printLog(msg string, req tReq, threads int) {
	fields := logseal.F{
		"url":      req.URL,
		"ipv4":     rxFindIPv4(req.Body),
		"ipv6":     rxFindIPv6(req.Body),
		"duration": req.Duration,
	}
	if CLI.LogLevel == string("trace") {
		fields["body"] = req.Body
		fields["threads"] = threads
	}
	if req.Error != nil {
		fields["error"] = req.Error
		lg.Error(msg, fields)
		return
	}
	if CLI.LogLevel == string("trace") {
		lg.Trace(msg, fields)
	}
	if CLI.LogLevel == string("debug") {
		fields["threads"] = threads
		lg.Debug(msg, fields)
	}
	if CLI.LogLevel == string("info") {
		lg.Info(msg, fields)
	}
}
