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

// similar to catch-and-throw assertations in java but the go equivalent that is super duper ghetto imo. by the way, nil is like null but it cannot be assigned to a variable.
func handleErr(err error){
	if err != nil {
		panic(err)
	}
}

func (s *SimpleServer) Address() string {return s.addr}

func (s *SimpleServer) isAlive() bool {return true}

func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request){
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) goNext() Server {
	server := lb.servers[lb.roundRobinCount % len(lb.servers)]
	for !server.isAlive(){
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount % len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// http package comming in clutch lmao
func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, req *http.Request){
	mainServer := lb.goNext()
	fmt.Printf("forwarding request to address %q\n", mainServer.Address())
	mainServer.Serve(rw, req)
}

// main function that starts everything up
func main() {
	servers := []Server{
		newSimpleServer("https://github.com/yehdar"),
		newSimpleServer("https://github.com/fakeshell"),
		newSimpleServer("https://github.com/huaanth"),
	}
	lb := newLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request){
		lb.serverProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)

	http.ListenAndServe(":" + lb.port, nil)
}