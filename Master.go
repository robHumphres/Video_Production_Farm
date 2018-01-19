package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

var connections []net.Conn

func MasterStartup() {
	// //all the parameters after master are ip addresses
	// ipAddresses := os.Args[2:]

	// fmt.Println(ipAddresses)
	go startMasterTCPServer()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload", UploadFile)
	log.Fatal(http.ListenAndServe(":9000", router))
}

func startMasterTCPServer() {

	ln, _ := net.Listen("tcp", ":27000")

	for {
		// listen on all interfaces

		// accept any connection on port 27000
		conn, _ := ln.Accept()
		connections = append(connections, conn)
		fmt.Println("recevied a connection")
		// run loop forever (or until ctrl-c)
		// will listen for message to process ending in newline (\n)
		// message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		// fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := "Connected to Master...."

		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
