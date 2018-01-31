# Video_Production_Farm
Video Production Server that is running software to control maya


To run as Master:</br>
GoLang: go run main.go Master.go ServiceUtils.go Slave.go Master</br>
Executable: Or ./executable Master

To Run as Slave:</br>
GoLang: go run main.go Master.go ServiceUtils.go Slave.go Slave host_ip:27000</br>
Executable: ./executable Slave host_ip:27000 something.py

Example Postman:

{

"file":"file.zip"

"scenes": "200"

}
</br>
How it works?</br>
This program is used as a TCP Master Slave bond and a REST-API for the ability for any user to use it. It does require the file being sent over to be in a .zip since the slaves need to receive it in that way.</br>

The Master which doesn't do any rendering can be hosted on any computer, and serves as a REST API and transport layer. It's job is to dynamically distribute the scenes based on the number sent over in the REST API. The Master can hold onto as many slaves as you want. It'll allocate the rendering process based on slaves stored in the list and divide them equally.</br>

Slaves will run in a state that they're always listening for master to get the order to render. They're expecting a zip file from master and what set of scenes to render, based off that they'll render those scenes. The rendering takes place in a python script that the user specifies as the 3rd arg when setting up the slave. After the specified scenes are render the Master which spins up a new thread for each connection will be listening for the files to come back and will put the folder separately from each other. All you'll have to do then is piece together the file.</br>
