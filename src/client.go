package main

import (
	"bufio"
	"bytes"
	"encoding/gob"

	"fmt"
	"net"
	"os"

	"github.com/es0j/topdownshooter2/util"
)

func main() {

	person := util.Person{Packetid: 1, Name: "John Doe", Age: 30}
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(person); err != nil {
		fmt.Println("Error encoding struct:", err)
		return
	}

	serializedData := b.Bytes()
	fmt.Println("Serialized data:", serializedData)

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send a message to the server
	_, err = conn.Write(serializedData)
	fmt.Println("send...")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is send
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data read from the connection to the terminal
	fmt.Print("> ", string(data))
}
