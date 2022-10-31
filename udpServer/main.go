package main

import (
	"log"
	"net"
)

const UDP_PROT string = ":8080"

func main() {
	StartUdpListen(UDP_PROT)
}

func StartUdpListen(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	udpConn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer udpConn.Close()
	log.Println("listener for server address=", addr)
	for {
		buf := make([]byte, 65507) //UDP最大理论长度65507
		n, addr, err := udpConn.ReadFromUDP(buf[0:])
		buf = buf[:n]
		if err != nil || n == 0 {
			log.Printf("%s %s %s, readLen=%d", addr.Network(), addr.String(), err.Error(), n)
			continue
		}
		if buf[n-1] != byte('\n') {
			buf = append(buf, '\n')
		}
		log.Printf("%s", buf)
	}
}
