package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"strconv"
	"os/exec"
	"bytes"
	b64 "encoding/base64"
	"os"
	"io/ioutil"
	"strings"
)


type NugArg struct {
	TheData []byte
	Inode string
}

type NugTSK struct {
	SavedData []byte
	PathToImage string
}

func (nd *NugTSK) GetDataLen(dataArg *NugArg, reply *string) error {
	*reply = strconv.Itoa(len(nd.SavedData))
	return nil
}

func (nd *NugTSK) LoadData(dataArg *NugArg, reply *string) error {
	if len(strings.Split(string(dataArg.TheData),":")) == 1 {
		fmt.Println("[-] Data location did not contain type. Assuming file.")
	}
	fmt.Print("[-] Loading data from:  " + string(dataArg.TheData))
	nd.SavedData = dataArg.TheData

	//we're going to hack it for now
	// load /space/m57/jo-2009-11-24.E01 into memory
	// actually.. no. just keep a reference ot it
	nd.PathToImage = string(dataArg.TheData)

	*reply = "done"
	fmt.Println(" .. done.")
	return nil
}

func (nd *NugTSK) ExecImageInfo(dataArg *NugArg, reply *string) error {
	//todo: log the hash of the tool used? at least get version information
	pathToTool := "/usr/bin/fsstat"
	cmd := exec.Command(pathToTool, "-i", "ewf", "-o", "63", nd.PathToImage)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	fmt.Println(out.String())
	*reply = out.String()

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (nd *NugTSK) GetBodyFile(dataArg *NugArg, reply *string) error {
	pathToTool := "/usr/bin/fls"

	fmt.Print("[-] Getting body file")
	cmd := exec.Command(pathToTool, "-rm", "/", "-o", "63", nd.PathToImage)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	//fmt.Println(out.String())
	*reply = out.String()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" .. done.")
	return err
}


func (nd *NugTSK) GetFileData(dataArg *NugArg, reply *string) error {
	pathToTool := "/usr/bin/icat"

	cmd := exec.Command(pathToTool, "-o", "63", nd.PathToImage,  dataArg.Inode)
	outfile, err := os.Create("tmpICATout")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start(); if err != nil {
		panic(err)
	}
	cmd.Wait()

	filedata, err := ioutil.ReadFile("tmpICATout")
	if err != nil {
		panic(err)
	}


	sEnc := b64.StdEncoding.EncodeToString(filedata)
	*reply = sEnc
	if err != nil {
		fmt.Println(err)
	}
	return err
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
