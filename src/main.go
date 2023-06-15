package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface{
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type SimpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &SimpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port 			string
	roundRobinCount int
	servers 		[]Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: 			 port,
		roundRobinCount: 0,
		servers:		 servers,
	}
}

func handleErr(err error){
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}