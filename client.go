package main

import (
	"log"
	"net/rpc"
)

type Client struct {
	client *rpc.Client
}

func (this *Client) Get(key string) string {
	var reply string
	err := this.client.Call("Server.Get", &Get{key}, &reply)
	if err != nil {
		log.Fatal("Get error:", err)
	}
	return reply
}

func (this *Client) Put(key string, val string) bool {
	var reply bool
	err := this.client.Call("Server.Put", &Put{key, val}, &reply)
	if err != nil {
		log.Fatal("Put error:", err)
	}
	return reply
}

func NewClient(addr string) *Client {
	// Make a client.
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	pc := new(Client)
	pc.client = client
	return pc
}
