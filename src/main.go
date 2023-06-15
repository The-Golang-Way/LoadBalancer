package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type SimpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port 			string
	roundRobinCount int
	servers			[]Server
}

func handleErr(err error){
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}