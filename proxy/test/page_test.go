/*
 * @Author: uyplayer
 * @Date: 2022/3/15 11:57 PM
 * @Email: uyplayer@qq.com
 * @File: page_test.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package test

import (
	"DDoSbyGolang/proxy"
	"sync"
	"testing"
)

func TestProxy(t *testing.T) {

	p := proxy.Proxy{
		Start:     2,
		End:       100,
		Wg:        &sync.WaitGroup{},
		Workernum: 50,
		Ch:        make(chan proxy.Workers, 50),
		Func:      proxy.HandleWorker,
		Proxy:     "161.35.161.38",
	}
	p.Workers()
	p.Generator()
	p.Wait()

}
