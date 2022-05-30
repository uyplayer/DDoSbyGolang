/*
 * @Author: uyplayer
 * @Date: 2022/3/16 10:59 PM
 * @Email: uyplayer@qq.com
 * @File: url.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package proxy

import "DDoSbyGolang/common"

//
//  GetUrl
//  @Description:解释json文件获取url
//  @receiver p
//  @return string
//
func (p *Proxy) GetUrl() string {
	// 	读出文件
	data, err := ReadJsonFile()
	if err != nil {
		common.Error.Fatalln(err)
	}
	// 获取url
	url := data.Proxyroviders[0].Url
	return url
}
