package simple

import (
	"fmt"
	"net"
)

func SimpleServer() {
	listenerConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 11111,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("local: %s \n", listenerConn.LocalAddr().String())

	// 读写缓冲区
	data := make([]byte, 1024)
	for {
		// 接收各个客户端发送的数据报
		n, remoteAddr, err := listenerConn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		// 写数据给特定的客户端
		_, err = listenerConn.WriteToUDP([]byte("honey-yogurt"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}
