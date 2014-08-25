package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	hashMap *HashMap
}

type Get struct {
	Key string
}

func (this *Server) Get(args *Get, reply *string) error {
	val, _ := this.hashMap.Get(args.Key)
	*reply = val.(string)
	return nil
}

type Put struct {
	Key, Val string
}

func (this *Server) Put(args *Put, reply *bool) error {
	this.hashMap.Put(args.Key, args.Val)
	*reply = true
	return nil
}

func NewServer(port string) {
	ps := new(Server)
	ps.hashMap = NewHashMap()
	rpc.RegisterName("Server", ps)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
