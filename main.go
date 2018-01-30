package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*
http://www.mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket/
https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

https://tour.golang.org/moretypes/15
https://golang.org/pkg/net/
https://awmanoj.github.io/tech/2016/12/16/keep-alive-http-requests-in-golang/
https://tour.golang.org/flowcontrol/12
http://www.mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket/
https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go


//Maya Stuff
https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-EB558BC0-5C2B-439C-9B00-F97BCB9688E4-htm.html
https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-1B7F8687-46C6-44CB-B224-C32A6B927AE8-htm.html

//Go MVC Stuff
https://github.com/utronframework/tutorials/blob/master/create_todo_list_application_with_utron.md
https://gobyexample.com/command-line-arguments
*/

var masterOrSlave = ""

//These buffer sizes MUST be the same on the slave as they are on the master (not sure if it matters since they're globals)
var bufferFileName = make([]byte, 64)
var bufferFileSize = make([]byte, 10)

const BUFFERSIZE = 1024

func main() {

	masterOrSlave := os.Args[1]
	if len(os.Args) > 1 {
		if masterOrSlave == "Master" {
			MasterStartup()
		} else {
			SlaveStartup(os.Args[2])
		}

	} else {
		fmt.Println("Error... Haven't supplied the right amount of argments\n" +
			"Args are as followed...\n" +
			"Master or Slave\n" +
			"If Master supply IP Addresses... as so 192.xxx.xx0 192.xxx.xx1 192.xxx.xx2\n" +
			"")
	}

}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("tcp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
