package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Embedding struct
type Packetid int32

type BasePacket struct {
	Packetid int32
}

type Person struct {
	Packetid
	Name string
	Age  int
}

const ( // iota is reset to 0
	p0 = iota // c0 == 0
	p1
	p2
)

/*func Decode(dec *gob.Decoder, e any) {
	if err := dec.Decode(&e); err != nil {
		fmt.Println("Error decoding struct:", err)
		return
	}
}*/

func Deserialize(received []byte) {
	var packet BasePacket
	dec := gob.NewDecoder(bytes.NewReader(received))

	//Decode(dec, packet)
	dec.Decode(&packet)
	fmt.Println("Deserialized struct:", packet)

	dec = gob.NewDecoder(bytes.NewReader(received))

	switch packet.Packetid {
	case p0:
		fmt.Println("p0")

	case p1:
		fmt.Println("p1")
		var p Person
		dec.Decode(&p)
		fmt.Println("person", p)

	}

}
