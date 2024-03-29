package main

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/triole/logseal"
)

type tReq struct {
	URL      string
	Body     string
	IPv4     string
	IPv6     string
	Duration time.Duration
	Error    error
}

func makeReq(url string, chin chan tReq, chout chan tReq) {
	chin <- tReq{URL: url}
	req := req(tReq{URL: url})
	time.Sleep(3 * time.Second)
	chout <- req
	<-chin
}

func req(req tReq) (resp tReq) {
	start := time.Now()
	resp.URL = req.URL

	var data []byte
	url, err := url.Parse(req.URL)
	if err != nil {
		// lg.IfErrError("can not parse url", logseal.F{"error": err})
		resp.Error = err
		return
	}

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		// lg.Error("can not init request", logseal.F{"error": err})
		resp.Error = err
		return
	}
	request.Header.Set("User-Agent", conf.UA)

	// lg.Debug("make request", logseal.F{"url": req.URL})
	response, err := client.Do(request)
	if err != nil {
		// lg.Error("request failed", logseal.F{"error": err})
		resp.Error = err
		return
	}

	if err == nil {
		data, err = io.ReadAll(response.Body)
		lg.IfErrError("unable to read request response", logseal.F{"error": err})
		if err != nil {
			resp.Error = err
			return
		}
		resp.Body = string(data)
		resp.IPv4 = rxFindIPv4(resp.Body)
		resp.IPv6 = rxFindIPv6(resp.Body)
	}
	resp.Duration = time.Since(start)
	return
}
