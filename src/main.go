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
// creates a simple server on a port
type SimpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

// think of this as the constructor class from java 
func newSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &SimpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

// creates the loadbalancer
type LoadBalancer struct {
	port 			string
	roundRobinCount int
	servers 		[]Server
}

// think of this as the constructor class from java 
func newLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: 			 port,
		roundRobinCount: 0,
		servers:		 servers,
	}
}

// similar to catch-and-throw assertations but super duper ghetto
func handleErr(err error){
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	servers := []Server{
		newSimpleServer("https://github.com/yehdar"),
		newSimpleServer("https://github.com/fakeshell"),
		newSimpleServer("https://github.com/huaanth"),
	}
	lb := newLoadBalancer("8000", servers)
}