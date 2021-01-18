package smt

import (
	"io"
	"log"
	"net"
	"os"
)

func MakeProof(fn func (io.Writer, io.Reader), port string) {
    conn, err := net.Dial("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fn(os.Stdout, conn)
}

func ClockServerProof(dst io.Writer, src io.Reader) {
	mustCopy(dst, src)
}

func EchoServerProof(dst io.Writer, src io.Reader) {
	go mustCopy(dst, src)
	mustCopy(src, dst)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
