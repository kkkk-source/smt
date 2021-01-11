// Clock is a TCP server that periodically writes the time.
package eg

import (
	"io"
	"log"
	"net"
	"time"
)

func ClockServer(port string) error {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// e.g, connection aborted
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
	return nil
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			// client disconnected
			return
		}
		time.Sleep(1 * time.Second)
	}
}
