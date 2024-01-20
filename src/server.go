package main

import (
	"fmt"
	"net"
	"os"

	"github.com/es0j/topdownshooter2/util"
)

func handleUDPClient(conn *net.UDPConn) {
	// Read from UDP listener in endless loop
	for {
		received := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(received)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("connected with:", addr)
		//fmt.Println("payload with:", received)

		util.Deserialize(received)

		// Write back the message over UPD
		//conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		//conn.WriteToUDP(serializedData, addr)
	}

}

func main() {

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("waiting for connection")
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connMap := make(map[string]chan []byte)
	for {
		

		received := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(received)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("connected with:", addr)
		val, ok := connMap[addr];
		if (ok) {
			//do something here
		}
		else{
			connMap[addr] = make(chan []byte)
			connMap[addr]
		}

		//fmt.Println("payload with:", received)

		util.Deserialize(received)

		// Write back the message over UPD
		//conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		//conn.WriteToUDP(serializedData, addr)
	}

}
