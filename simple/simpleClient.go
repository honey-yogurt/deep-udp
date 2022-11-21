package simple

import (
	"fmt"
	"net"
)

func SimpleClient() {
	ip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}
	dstAddr := &net.UDPAddr{
		IP:   ip,
		Port: 11111,
	}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello"))
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		return
	}

	fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
}
