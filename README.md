# Video_Production_Farm
Video Production Server that is running software to control maya


To run as Master:</br>
GoLang: go run main.go Master.go ServiceUtils.go Slave.go Master</br>
Executable: Or ./executable Master

To Run as Slave:</br>
GoLang: go run main.go Master.go ServiceUtils.go Slave.go Slave host_ip:27000</br>
Executable: ./executable Slave host_ip:27000
