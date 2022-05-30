/*
 * @Author: uyplayer
 * @Date: 2022/3/29 23:31
 * @Email: uyplayer@qq.com
 * @File: echoserver.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / bhg
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package bhg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echoecho(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func echomain() {

	// bind server
	listener, err := net.Listen("tcp", "20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echoecho(conn)

	}
}

func echobuffer(conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes:%s", len(s), s)
	log.Println("Writing data")

	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString(s)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	// 写到里面
	writer.Flush()
}

type Work interface {
	yagaqqi()
	tomuqi()
}

type sarang interface {
	yagaqqi()
}

type sananti interface {
	aaaaa()
}
type aa struct {
}

func makedo(a sarang) {
	fmt.Println(" test ")

}

func check() {
	var s Work
	//var b sananti
	makedo(s)
	//makedo(b)

}
