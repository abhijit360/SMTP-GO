package main

import (
	// "bufio"
	"fmt"
	"net"
	"strings"
)


func main(){
	fmt.Println("hello from go")
	listener, err := net.Listen("tcp",":25")
	if err != nil {
		fmt.Println("Error accepting request:", err)
	}

	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occured during continued connection",err)
			return
		}
		handleMailConnection(conn)
	}

}

func handleMailConnection(conn net.Conn){
	fmt.Printf("connecting from: %v\n",conn.RemoteAddr())
	STATUS_OK := []byte("250 ok")
	buffer := make([]byte,1024)
	conn.Read(buffer)
	split_strings := strings.Split(string(buffer)," ")
	cmd := strings.ToLower(split_strings[0])

	switch(cmd){
	case "helo":
		// clear all the storage etc and return a 250 ok
		conn.Write(STATUS_OK)
	case "mail":
		defer conn.Close()
	}

}	