package tcp

import (
	"net"
	"log"
	"time"
)


func TestTcpConnection(service string, timeout int)(status bool){
	status=true
	test := net.Dialer{Timeout: time.Second*time.Duration(timeout)}
	log.Println("TCP Connect to : ", service)
	conn, err := test.Dial("tcp", service)
	if err != nil {
		status=false
		return status
	}
	defer conn.Close()
	return status
}