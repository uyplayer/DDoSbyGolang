/*
 * @Author: uyplayer
 * @Date: 2022/3/26 13:29
 * @Email: uyplayer@qq.com
 * @File: proxy.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / layer7
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package layer7

import (
	"DDoSbyGolang/common"
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

//
//  getProxy
//  @Description: 返回useragent
//  @receiver f
//
func GetProxystring() string {

	// 形成随机数
	rand.Seed(time.Now().UnixNano())
	// 打开useragent文件
	openFile, err := os.Open("/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/proxyIP.txt")
	defer openFile.Close()
	if err != nil {
		common.Error.Fatalln(err)
	}

	// 需要一行行读出
	var proxyList []string
	bf := bufio.NewReader(openFile)
	for {
		// 返回新的读出的对象
		line, _, err := bf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		// 添加
		proxyList = append(proxyList, string(line))
	}
	// 随机数
	any := rand.Intn(len(proxyList))
	return proxyList[any]

}
