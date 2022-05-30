/*
 * @Author: uyplayer
 * @Date: 2022/3/10 3:12 PM
 * @Email: uyplayer@qq.com
 * @File: logs.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package common

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//  初始化log变量
var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

//
//  init
//  @Description: 初始化log
//
func init() {

	// error 写入文件
	file, err := os.OpenFile("/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func m() {
}
