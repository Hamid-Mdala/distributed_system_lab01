package main

import (
	"log"
	"net"
	"net/rpc"
	"strings"
)

type Args struct {
	Text string
}

type Count struct {
	Count int
}

type Service struct{}

func (s *Service) Count(args *Args, reply *Count) error {
	reply.Count = len(strings.Fields(args.Text)) // counts words
	return nil
}

func main() {
	svc := new(Service)
	rpc.Register(svc)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}
	log.Println("RPC server running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
