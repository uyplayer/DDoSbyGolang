/*
 * @Author: uyplayer
 * @Date: 2022/3/11 9:40 PM
 * @Email: uyplayer@qq.com
 * @File: page.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

// https://www.jianshu.com/p/1115f3c71a18

package proxy

import (
	"fmt"
	"math/rand"
	"time"
)

//
//  workers
//  @Description: pool 池
//  @receiver p
//
func (p *Proxy) Workers() {
	fmt.Println(" Workers 启动 ")
	// 获取user agent
	userAgent := ProcessUserAgent()
	// 创建work
	for i := 0; i < p.Workernum; i++ {
		// 添加
		p.Wg.Add(1)
		// 开启新的go routine
		go func(workId int) {
			defer p.Wg.Done()
			// 通道等待信息
			for work := range p.Ch {
				// 调用处理函数
				// 随机返回一个useragent
				// 还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子
				// time.Now().UnixNano(): 当前操作系统时间的毫秒值
				rand.Seed(time.Now().UnixNano())
				arbNum := rand.Intn(len(userAgent))
				ua := userAgent[arbNum]
				p.Func(&work, workId, ua)
			}
		}(i)

	}

}

//
//  generator
//  @Description: 形成url
//  @receiver p
//
func (p *Proxy) Generator() {
	fmt.Println(" Generator 启动 ")
	// 获取url
	url := p.GetUrl()
	// 开始页
	start := p.Start
	// 接触页面
	end := p.End
	// 新goroutine运行
	go func() {
		// 形成rul
		for ; start < end; start++ {
			// 格式化url
			url := fmt.Sprintf("%s/%d.html", url, start)
			// 创建work
			work := Workers{
				url: url,
			}
			// 对象传到worker
			p.Ch <- work
		}
		// 关闭channel
		close(p.Ch)
	}()
}

//
//  wait
//  @Description: 等待结束
//  @receiver p
//
func (p *Proxy) Wait() {
	fmt.Println(" Wait 启动 ")
	// 等待结束
	p.Wg.Wait()

}

//
//  InitProxyPool
//  @Description: 初始化
//  @receiver p
//  @return *Proxy
//
//func (p *Proxy) InitProxyPool() *Proxy {
//	return &Proxy{
//		Start:     2,
//		End:       100,
//		Wg:        &sync.WaitGroup{},
//		channum:   50,
//		workernum: 50,
//		ch:        make(chan Workers, channum),
//		Func:      handleWorker,
//		Proxy:     "203.76.117.18",
//	}
//}
