package main

import (
	"fmt"
	"net"
	"os"
)

func SlaveStartup(masterIP string) {
	fmt.Println("It's a slave")

	//don't like seeing the red line
	fmt.Println("Master ip in slave:" + masterIP)

	if len(os.Args) < 2 {
		fmt.Println("Not enough Arguments are supplied... Must specific Master IP")
		os.Exit(-1)
	}

	fmt.Println("Start tcp slace with : " + masterIP)

	StartTCPConnection(masterIP)

}

func StartTCPConnection(masterIP string) {

	fmt.Println("slave connecting too... " + masterIP)
	connection, err := net.Dial("tcp", masterIP)

	if err != nil {
		fmt.Println("Master is not online....")
		panic(err)
	}
	fmt.Println("Connected to master.... Now waiting for files to render\n")

	defer connection.Close()

	//4Eva Loop for the TCP Reading and writing
	for {
		//Waits for server to send stuff
		connection.Read(bufferFileSize)

		// fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

		connection.Read(bufferFileName)
		// fileName := strings.Trim(string(bufferFileName), ":")

		fmt.Println("From Server... : " + string(bufferFileName))

		// newFile, err := os.Create(fileName)

		// if err != nil {
		// 	panic(err)
		// }

		// defer newFile.Close()
		// var receivedBytes int64

		// for {
		// 	if (fileSize - receivedBytes) < BUFFERSIZE {
		// 		io.CopyN(newFile, connection, (fileSize - receivedBytes))
		// 		connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
		// 		break
		// 	}
		// 	io.CopyN(newFile, connection, BUFFERSIZE)
		// 	receivedBytes += BUFFERSIZE
		// }
		// fmt.Println("Received file completely!")

		// prepareRendering(fileName)
		// startRendering(fileName)
		// sendRendering(fileName)
	}

}
