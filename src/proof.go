package smt

import (
	"io"
	"log"
	"net"
	"os"
)

// MakeProof make easy to create a ClockServerProof and a EchoServerProof.
func MakeProof(fn func(io.ReadWriter), port string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fn(conn)
}

// ClockServerProof is a TCP read-only client. You can
// use it to read from the ClockServer output.
func ClockServerProof(src io.ReadWriter) {
	mustCopy(os.Stdout, src)
}

// EchoServerProof is TCP read-write client. You can
// use it to read from and write to the EchoServer.
func EchoServerProof(src io.ReadWriter) {
	go mustCopy(os.Stdout, src)
	mustCopy(src, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
