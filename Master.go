package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var connections []net.Conn
var fileNamePOST string
var sendBuffer []byte

//MasterStartup Starts the Master TCP Server and REST API
func MasterStartup() {
	// //all the parameters after master are ip addresses
	// ipAddresses := os.Args[2:]

	// fmt.Println(ipAddresses)
	go startMasterTCPServer()

	//Start Rest Server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload", UploadFile)
	log.Fatal(http.ListenAndServe(":9000", router))
}

//startMasterTCPServer starts the tcp listener for slaves to connect to
func startMasterTCPServer() {

	server, _ := net.Listen("tcp", ":27000")

	defer server.Close()

	for {

		// accept any connection on port 27000
		conn, err := server.Accept()

		if err != nil {
			panic(err)
		}

		connections = append(connections, conn)
		fmt.Println("Received a connection from.. " + conn.RemoteAddr().String())
	}
}

//UploadFile REST API Function that only allows POST Methods, Then send out to slaves for rendering
func UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recieved POST")
	file, handler, err := r.FormFile("file")
	defer file.Close()

	fmt.Println("File name is... " + handler.Filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Fprintf(w, "%v", handler.Header)
	fmt.Fprintf(w, "Received...")

	//Create the file
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	io.Copy(f, file)

	fileNamePOST = handler.Filename

	PostToSlaves()

}

//PostToSlaves is what's used to post to multiple slaves for rendering *Needs to figure out rendering*
func PostToSlaves() {

	//Needs some time for the file to set from post
	time.Sleep(5 * time.Second)

	for i := range connections {

		//Transfer the name
		fmt.Println("Made it to connection... " + string(i))

		file, err := os.Open(fileNamePOST)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}

		fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
		fileName := fillString("1"+fileInfo.Name(), 64)
		fmt.Println("Sending filename and filesize!")
		connections[i].Write([]byte(fileSize))
		connections[i].Write([]byte(fileName))
		sendBuffer := make([]byte, BUFFERSIZE)
		fmt.Println("Start sending file!")

		//Sending file
		for {
			_, err = file.Read(sendBuffer)
			if err == io.EOF {
				break
			}
			connections[i].Write(sendBuffer)
		}

		fmt.Println("\nFinished sending files to slaves")

		//New thread to each slave
		go ReceiveFileFromSlave(connections[i])

	}
}

//ReceiveFromSlave starts on a new thread per client
func ReceiveFileFromSlave(conn net.Conn) {
	var buffer []byte
	conn.Read(buffer)

	fmt.Println(strings.Trim(string(buffer), ":"))

	fmt.Println("Received file from.... " + conn.RemoteAddr().String())

}
