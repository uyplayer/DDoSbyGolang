/*
 * @Author: uyplayer
 * @Date: 2022/3/30 20:53
 * @Email: uyplayer@qq.com
 * @File: tcpproxy.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / bhg
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package bhg

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	// Run in goroutine to prevent io.Copy fromblocking ❷
	go func() {
		// Copy oursource's outputto the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	//Copyourdestination'soutputbacktooursource
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

func server() {

	//Listenonlocalport80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unabletoacceptconnection")
		}
		go handle(conn)
	}
}
