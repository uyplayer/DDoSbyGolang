/*
 * @Author: uyplayer
 * @Date: 2022/3/26 13:24
 * @Email: uyplayer@qq.com
 * @File: agents.go
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
//  getGoogleAgents
//  @Description: 随机获取Google agent
//  @return string
//
func getGoogleAgents() string {

	// 形成随机数
	rand.Seed(time.Now().UnixNano())

	// 谷歌代理
	googleAgents := []string{
		"Mozila/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; http://www.google.com/bot.html)) Googlebot/2.1 (+http://www.google.com/bot.html)",
		"Googlebot/2.1 (+http://www.googlebot.com/bot.html)",
	}
	// 随机数
	any := rand.Intn(len(googleAgents))

	return googleAgents[any]

}

//
//  getUserAgent
//  @Description: 返回useragent
//  @receiver f
//
func getUserAgent() string {
	// 形成随机数
	rand.Seed(time.Now().UnixNano())
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

		userAgentList = append(userAgentList, string(line))
	}
	// 随机数
	any := rand.Intn(len(userAgentList))
	return userAgentList[any]

}
