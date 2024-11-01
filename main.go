package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"io"
)

var connectionCodes = map[string]int{
    "STATUS_OK":                    250,
    "TRANSACTION_FAILED":            554,
    "SERVICE_READY":                 220,
    "SERVICE_CLOSING":               221,
    "SERVICE_NOT_AVAILABLE":         421,
    "MAILBOX_NOT_AVAILABLE":         450, // same for mailbox not available
    "REQUESTED_ACTION_NOT_TAKEN":    550,
}

type email struct{
	from string
	to string
	content string
}

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
		defer listener.Close()
	}

}

func formatMessage(status_code int , message string) []byte {
	return []byte(fmt.Sprintf("%d %s\r\n", status_code, message))
}

func handleMailConnection(conn net.Conn){
	fmt.Printf("connecting from: %v\n",conn.RemoteAddr())
	buffer := make([]byte,1024)
	_ , err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:",err)
		return
	}
	split_strings := strings.Split(string(buffer)," ")
	cmd := strings.ToLower(split_strings[0])
	
	currentEmail := email{"","",""}

	switch(cmd){
	case "helo":
		// clear all the storage etc and return a 250 ok		
		conn.Write(formatMessage(connectionCodes["STATUS_OK"],"Ready to get email!"))
	case "mail":
		domain := split_strings[1]
		currentEmail.from = string(domain)

		conn.Write(formatMessage(connectionCodes["STATUS_OK"],"ok"))
	case "data":
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString("\n");
			if err == io.EOF{
				break
			}
			if strings.TrimSpace(line) == "."{
				break
			}
			currentEmail.content = fmt.Sprintf("%v\n %v", currentEmail.content,line)
		}
		conn.Write(formatMessage(connectionCodes["STATUS_OK"], "Message received"))
	}

}	