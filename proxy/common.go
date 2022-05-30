/*
 * @Author: uyplayer
 * @Date: 2022/3/11 9:39 PM
 * @Email: uyplayer@qq.com
 * @File: pool.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package proxy

import "sync"

// proxy 提供商
const file string = "/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/proxy-providers.json"
const proxyIP = "/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/proxyIP.txt"

//  处理器
type handler func(workers *Workers, int2 int, ua string)

//
//  Workers
//  @Description: workers
//
type Workers struct {
	url string
}

//
//  Proxy
//  @Description: proxy对象
//
type Proxy struct {
	Start     int             `json:"start"`
	End       int             `json:"end"`
	FileLoc   string          `json:"file_loc"`
	Wg        *sync.WaitGroup `json:"wg"`
	Channum   int             `json:"channum"`
	Workernum int             `json:"workernum"`
	Ch        chan Workers    `json:"ch"`
	Func      handler         `json:"func"`
	Proxy     string          `json:"proxy"`
}
