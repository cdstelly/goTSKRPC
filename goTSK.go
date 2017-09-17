package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
	"net/http"
)

type NugArg struct {
	TheData []byte
}

type NugTSK struct {
	SavedData []byte
}

func (nd *NugTSK) LoadDataTSK(dataArg *NugArg, reply *string) error {
	nd.SavedData = dataArg.TheData
	*reply = "done"
	return nil
}

func main() {
	fmt.Println("started")
	tsk := new(NugTSK)
	rpc.Register(tsk)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":2001")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	http.Serve(l,nil) //won't pass here without an error
	fmt.Println("done")
}
