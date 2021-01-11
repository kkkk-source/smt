package eg

import (
    "io"
    "fmt"
    "log"
    "net"
    "time"
    "bufio"
    "strings"
)

// localhost = 0.0.0.0
const ip = "0.0.0.0"

// MakeServer make easy to create this package's
// servers, e.g, ClockServer or EchoServer.
func MakeServer(fn func(net.Conn), port string) error {
    if port == "" {
        port = "8000"
    }
    listener, err := net.Listen("tcp", ip+":"+port)
    if err != nil {
        return err
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            // connection aborted
            log.Print(err)
            continue
        }
        go fn(conn)
    }
    return nil
}

// ClockServer is a TCP server that periodically writes the time.
func ClockServer(c net.Conn) {
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

// EchoServe Simulate the reverberations of a real echo, with 
// the response loud at first ("WTF!"), then moderate ("Wtf!") 
// after a delay, then quiet ("wft!") before fading to nothing.
func EchoServer(c net.Conn) {
    input := bufio.NewScanner(c)
    for input.Scan() {
        echo(c, input.Text(), 1*time.Second)
    }
}

func echo(c net.Conn, shout string, delay time.Duration) {
    shout = strings.ToLower(shout)
    fmt.Fprintln(c, "\t>> ", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t>> ", strings.Title(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t>> ", shout)
}
