// eg package provide a variaty of concurrent programs that you can
// use to play with. becasue Networking is a natural domain in which
// to use concurrency since servers typically handle many connections
// from their clients at once, each client being essentially independent
// of the others, there are some servers applications as well.
package eg

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DiskUsage computes the disk usage of the files in a directory.
func DiskUsage(roots []string) {
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// SpinnerAnimation computes the 45th Fibonacci number. Since it
// uses the terribly inefficient recursive algorithm, it runs for
// an appreciable time, during which it provide the user with a
// visual indication that the program is still running by displaying
// an animated textual "spinner".
func SpinnerAnimation() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fibonacci(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// localhost is equals to 0.0.0.0
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
