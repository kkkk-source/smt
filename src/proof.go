package smt

import (
	"io"
	"log"
	"net"
	"os"
)

// MakeProof make easy to create a ClockServerProof and a EchoServerProof.
func MakeProof(fn func(io.Writer, io.Reader), port string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fn(os.Stdout, conn)
}

// ClockServerProof is a TCP read-only client. You can 
// use it to read from the ClockServer output.
func ClockServerProof(dst io.Writer, src io.Reader) {
	mustCopy(dst, src)
}

// EchoServerProof is TCP read-write client. You can 
// use it to read from and write to the EchoServer.
func EchoServerProof(dst io.Writer, src io.Reader) {
	go mustCopy(dst, src)
	mustCopy(src, dst)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
