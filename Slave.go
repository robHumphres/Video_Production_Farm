package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var pythonScriptName string
var masterConnection net.Conn

// SlaveStartup basic function tells what's the master ip you set, and the python script being used starts TCP Connection
func SlaveStartup(masterIP string, pythonScript string) {
	fmt.Println("It's a Slave!")

	fmt.Println("Master ip in slave: " + masterIP)
	fmt.Println("Python script being used... " + pythonScript)

	pythonScriptName = pythonScript
	StartTCPConnection(masterIP)

}

// StartTCPConnection Tries to connect to Master, Waits for Master to Send out File, and then calls Render Function
func StartTCPConnection(masterIP string) {

	connection, err := net.Dial("tcp", masterIP)

	if err != nil {
		fmt.Println("Master is not online....")
		panic(err)
	}
	fmt.Println("Connected to master.... Now waiting for files to render\n")

	defer connection.Close()

	//4Eva Loop for the TCP Reading and writing
	for {
		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")

		newFile, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}
		defer newFile.Close()
		var receivedBytes int64

		for {

			if (fileSize - receivedBytes) < BUFFERSIZE {
				io.CopyN(newFile, connection, (fileSize - receivedBytes))
				connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
				break
			}
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}
		fmt.Println("Received file completely ! " + fileName)
		// renderFile()
	}

}

// renderFile takes the file given from the Master, calls script with parameters and then calls SendRenderedFile
func renderFile() {
	/*
		- Figure what flags are needed for maya to render
		- Call Global Var for what script to run

	*/

	//TODO: Need to know what
	//Calls Script
	cmd := exec.Command(pythonScriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}

func sendRenderedFile() {
	masterConnection.Write([]byte("something random"))
}
