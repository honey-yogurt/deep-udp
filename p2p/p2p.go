package p2p

import (
	"fmt"
	"net"
	"os"
	"time"
)

func read(conn *net.UDPConn) {
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err.Error())
		}
		fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}

func P2p() {
	addr1 := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 11111,
	}
	addr2 := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 11112,
	}

	go func() {
		listenConn1, err := net.ListenUDP("udp", addr1)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listenConn1)
		time.Sleep(1 * time.Second)
		listenConn1.WriteToUDP([]byte("ping to #2: "+addr2.String()), addr2)
	}()

	go func() {
		listenConn2, err := net.ListenUDP("udp", addr2)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listenConn2)
		time.Sleep(1 * time.Second)
		listenConn2.WriteToUDP([]byte("ping to #2: "+addr1.String()), addr1)
	}()
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
