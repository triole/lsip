package main

import (
	"github.com/triole/logseal"
)

type tDone struct {
	IPv4 bool
	IPv6 bool
	Done bool
}

func process() {
	chin := make(chan tReq, CLI.Threads)
	chout := make(chan tReq, CLI.Threads)

	for _, url := range conf.IPRetrieval {
		if CLI.DryRun {
			lg.Info(
				"dry run, would have made a request",
				logseal.F{"url": url})
		} else {
			go makeReq(url, chin, chout)
		}
	}

	if !CLI.DryRun {
		c := 0
		ln := len(conf.IPRetrieval)
		done := tDone{IPv4: false, IPv6: false, Done: false}
		for req := range chout {
			c++
			// default mode, print first fetched ip matching the grep flag
			if req.Error == nil {
				if CLI.Print == "4" && isValidIPv4((req.IPv4)) {
					printLog("successful ipv4 fetch", req, len(chin))
					done.IPv4 = true
					done.Done = true
				}

				if CLI.Print == "6" && isValidIPv6((req.IPv6)) {
					printLog("successful ipv6 fetch", req, len(chin))
					done.IPv6 = true
					done.Done = true
				}

				if CLI.Print == "both" && (isValidIPv4(req.IPv4) || isValidIPv6(req.IPv6)) {
					if isValidIPv4(req.IPv4) && !done.IPv4 {
						printLog("successful ipv4 fetch, both ip versions", req, len(chin))
						done.IPv4 = true
					}
					if isValidIPv6(req.IPv6) && !done.IPv6 {
						printLog("successful ipv6 fetch, both ip versions", req, len(chin))
						done.IPv6 = true
					}
					if done.IPv4 && done.IPv6 {
						done.Done = true
					}
				}

				if CLI.Print == "any" && (isValidIPv4(req.IPv4) || isValidIPv6(req.IPv6)) {
					printLog("successful fetch, any ip version", req, len(chin))
					done.Done = true
				}

				if CLI.Print == "all" && (isValidIPv4(req.IPv4) || isValidIPv6(req.IPv6)) {
					printLog("successful fetch, any ip version", req, len(chin))
				}

				if done.Done {
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
