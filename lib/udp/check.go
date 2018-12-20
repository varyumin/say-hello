package udp

import (
	"net"
	"time"
	"log"
)

func TestUdpConnection(service string, timeout int) (status bool){
	status=true
	log.Println("UDP Resolve: ", service)
	RemoteAddr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		status=false
		return status
		log.Println(err)
	}

	log.Println("UDP Connect to: ", RemoteAddr)
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		status=false
		return status
		log.Println(err)
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
	conn.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(timeout)))

	defer conn.Close()

	message := []byte(" HELLO!")
	log.Println("UDP Write message: ", message)
	_, err = conn.Write(message)
	if err != nil {
		status=false
		return status
		log.Println(err)
	}

	buffer := make([]byte, 1024)
	log.Println("UDP Read message... ")
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		status=false
		return status
		log.Println(err)
	}
	log.Println("UDP Server : ", addr)
	log.Println("Received from UDP server : ", string(buffer[:n]))

	return status
}
