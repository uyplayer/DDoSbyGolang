/*
 * @Author: uyplayer
 * @Date: 2022/3/26 13:44
 * @Email: uyplayer@qq.com
 * @File: home.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / layer7
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package layer7

import (
	"net"
	"sync"
)
import "context"

//  第七层协议
var LAYER7_METHODS []string = []string{
	"CFB", "BYPASS", "GET", "POST", "OVH", "STRESS", "DYN", "SLOW", "HEAD",
	"NULL", "COOKIE", "PPS", "EVEN", "GSB", "DGB", "AVB", "CFBUAM",
	"APACHE", "XMLRPC", "BOT", "BOMB", "DOWNLOADER",
}

//  干活的函数
type DoWork func(work Work)
type Panic func()

//
//  Work
//  @Description: worker
//
type Work struct {
	Target    string `json:"target"`
	Proxy     string `json:"Proxy"`
	UserAgent string `json:"user_agent"`
	Context   context.Context
	Port      string       `json:"port"`
	Url       string       `json:"url"`
	Method    string       `json:"method"`
	PanRec    Panic        `json:"pan_rec"`
	Con       *net.TCPConn `json:"con"`
	Payload   chan string  `json:"payload"`
}

//
//  Layer7
//  @Description: layer7
//
type Layer7 struct {
	Target       string          `json:"target"`
	Proxy        string          `json:"Proxy"`
	MaxWorkerNum int             `json:"max_worker_num"`
	Wg           *sync.WaitGroup `json:"wg"`
	Data         chan Work       `json:"data"`
	Func         DoWork          `json:"func"`
	Context      context.Context `json:"context"`
	Port         string          `json:"port"`
	Url          string          `json:"url"`
	Method       string          `json:"method"`
	Payload      chan string     `json:"payload"`
	RunTime      int             `json:"run_time"`
}
