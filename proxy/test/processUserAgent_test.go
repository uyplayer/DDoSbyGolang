/*
 * @Author: uyplayer
 * @Date: 2022/3/15 11:57 PM
 * @Email: uyplayer@qq.com
 * @File: processUserAgent_test.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package test

import (
	"DDoSbyGolang/common"
	"DDoSbyGolang/proxy"
	"testing"
)

func TestUserAgent(t *testing.T) {

	ua := proxy.ProcessUserAgent()
	if len(ua) == 0 {

		t.Error(" ops ")

	}
	common.Info.Println(ua)

}
