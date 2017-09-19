package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"strconv"
)

type NugArg struct {
	TheData []byte
}

type NugTSK struct {
	SavedData []byte
}

func (nd *NugTSK) GetDataLen(dataArg *NugArg, reply *string) error {
	*reply = strconv.Itoa(len(nd.SavedData))
	return nil
}

func (nd *NugTSK) LoadData(dataArg *NugArg, reply *string) error {
	nd.SavedData = dataArg.TheData
	*reply = "done"
	return nil
}

func (nd *NugTSK) ExecInfo(dataArg *NugArg, reply *string) error {
	*reply = "test"
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
