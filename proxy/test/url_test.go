/*
 * @Author: uyplayer
 * @Date: 2022/3/16 11:00 PM
 * @Email: uyplayer@qq.com
 * @File: url_test.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package test

import (
	"DDoSbyGolang/proxy"
	"log"
	"testing"
)

func TestUrl(t *testing.T) {

	p := proxy.Proxy{}
	if p.GetUrl() != "" {
		log.Fatal(" ops ")
	}

}
