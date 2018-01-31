# Video_Production_Farm
Video Production Server that is running software to control maya


To run as Master:
go run main.go Master.go ServiceUtils.go Slave.go Master</br>
Or ./<executable> Master

To Run as Slave:
go run main.go Master.go ServiceUtils.go Slave.go Slave <hostip>:27000</br>
or ./<executable> Slave <hostip>:27000
