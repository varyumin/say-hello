package icmp

import (
	"log"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"strings"
)


func TestTcpConnection(service string) (status bool) {
	s := strings.Split(service, ":")
	targetIP := s[0]
	c, err := icmp.ListenPacket("ip4:icmp", "127.0.0.1")
	if err != nil {
		log.Fatalf("listen err, %s", err)
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetIP)}); err != nil {
		log.Fatalf("WriteTo err, %s", err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	rm, err := icmp.ParseMessage(58, rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v", peer)
	default:
		log.Printf("got %+v; want echo reply", rm)
	}

	status=true
	return status
}