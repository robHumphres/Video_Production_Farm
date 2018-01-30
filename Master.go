package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var connections []net.Conn

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

func UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recieved request from postman")
	file, handler, err := r.FormFile("file")
	fmt.Printf("Post came through\n")
	defer file.Close()

	fmt.Println("File name is... " + handler.Filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Fprintf(w, "%v", handler.Header)
	fmt.Fprintf(w, "Received...")

	PostToSlaves(handler.Filename)

}

func PostToSlaves(fileame string) {

	var number int

	for i := range connections {

		//Transfer the name
		fmt.Println("Made it to connection... " + string(i))

		file, err := os.Open("exerciseTest.csv")
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
		fileName := fillString(string(number)+fileInfo.Name(), 64)
		fmt.Println("Sending filename and filesize!")
		connections[i].Write([]byte(fileSize))
		connections[i].Write([]byte(fileName))
		sendBuffer := make([]byte, BUFFERSIZE)
		fmt.Println("Start sending file!")
		for {
			_, err = file.Read(sendBuffer)
			if err == io.EOF {
				break
			}
			connections[i].Write(sendBuffer)
		}
		number++

	}
}

// //Post uploads a single file
// if r.Method == "POST" {
// 	file, handler, err := r.FormFile("file")
// 	fmt.Printf("Post came through\n")
// 	defer file.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Fprintf(w, "%v", handler.Header)
// 	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer f.Close()
// 	io.Copy(f, file)

// 	//Delete older ones if past 10
// 	UnzipNClean(handler.Filename)

// } else {
// 	fmt.Fprintf(w, "This is just a POST Method, see documentation")
// }

// return
