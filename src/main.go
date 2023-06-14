package main

import (
	"net/http/httputil"
	"net/url"
)

type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, bruhWhatDoIDoHere := url.Parse(addr)

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}