package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	Text string
}

type Count struct {
	Count int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Close() // good practice to close the connection

	args := Args{
		Text: "andy is here ad andy is strong",
	}

	var reply Count

	err = client.Call("Service.Count", &args, &reply)
	if err != nil {
		log.Fatal("RPC call error:", err)
	}

	// Fix: use Printf (with a format specifier) or Println
	fmt.Printf("Word count: %d\n", reply.Count)
	// Alternatively: fmt.Println("Word count:", reply.Count)
}
