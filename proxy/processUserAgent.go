/*
 * @Author: uyplayer
 * @Date: 2022/3/11 9:53 PM
 * @Email: uyplayer@qq.com
 * @File: processUserAgent.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package proxy

import (
	"DDoSbyGolang/common"
	"bufio"
	"io"
	"os"
)

//
//  ProcessUserAgent
//  @Description: 每次调用返回一个useragent
//  @receiver f
//
func ProcessUserAgent() []string {

	// 打开useragent文件
	openFile, err := os.Open("/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/useragent.txt")
	defer openFile.Close()
	if err != nil {
		common.Error.Fatalln(err)
	}

	// 需要一行行读出
	var userAgentList []string
	bf := bufio.NewReader(openFile)
	for {
		// 返回新的读出的对象
		line, _, err := bf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		//common.Info.Println(string(line))
		// 添加
		userAgentList = append(userAgentList, string(line))
	}

	return userAgentList

}
