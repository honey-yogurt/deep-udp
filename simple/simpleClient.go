package simple

import (
	"fmt"
	"net"
	"os"
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

	// 保持UDP Socket文件一直打开
	b := make([]byte, 1)
	os.Stdin.Read(b)

	conn.Write([]byte("hello"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())
}
