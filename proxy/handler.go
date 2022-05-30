/*
 * @Author: uyplayer
 * @Date: 2022/3/16 11:04 PM
 * @Email: uyplayer@qq.com
 * @File: handler.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package proxy

import (
	"DDoSbyGolang/common"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mahonia"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//
//  HandleWorker
//  @Description: 请求并解释
//  @param workers
//  @param int2
//
func HandleWorker(workers *Workers, int2 int, ua string) {
	common.Info.Printf(" worker id : %d ", int2)
	// 创建新的请求
	request, err := http.NewRequest("GET", workers.url, nil)
	// 添加user agent
	// 页面编码
	request.Header.Add("Content", "text/html; charset=gbk")
	request.Header.Add("Host", "www.66ip.cn")
	request.Header.Add("Referer", "http://www.66ip.cn/index.html")
	request.Header.Add("User-Agent", ua)
	if err != nil {
		common.Error.Fatalln(err)
	}
	client := &http.Client{}
	// 发送请求
	rep, err := client.Do(request)
	if err != nil {
		common.Error.Fatalln(err)
	}
	// 关闭请求
	defer rep.Body.Close()
	// 处理gbk编码
	buf := new(bytes.Buffer)
	buf.ReadFrom(rep.Body)
	s := buf.String()
	// gbk -》 utf-8
	result := gbkHandler(s, "gbk", "utf-8")
	//fmt.Println(result)
	// dom 解释
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		common.Error.Fatalln(err)
	}
	// 写入的文件
	f, err := os.OpenFile(proxyIP, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		common.Error.Fatalln(err.Error())
	}
	// dom 查询
	//fmt.Println(" 开始解释dom")
	dom.Find("div[align=center]>table>tbody>tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			ip := selection.Find("td").First().Text()
			port := selection.Find("td").First().Next().Text()
			fullip := fmt.Sprintf("%s:%s", ip, port)
			if checkIP(fullip) {
				f.Write([]byte(fullip + "\n"))
			}

		}

	})
}

//
//  gbkHandler
//  @Description: 处理gbk
//  @param src
//  @param srcCode
//  @param tagCode
//  @return string
//
func gbkHandler(src string, sourceCode, targetCode string) string {

	srcCoder := mahonia.NewDecoder(sourceCode)
	// 转换
	srcResult := srcCoder.ConvertString(src)
	// 转换
	tagCoder := mahonia.NewDecoder(targetCode)
	// 转换
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

//
//  checkIP
//  @Description: 查看ip是否有效
//  @param ip
//
func checkIP(ip string) bool {
	// 目标
	target := "http://www.baidu.com"
	// 定义Transport
	tr := &http.Transport{}
	tr.DisableKeepAlives = true
	if len(ip) != 0 { // Set the proxy only if the proxy param is specified
		proxyUrl, err := url.Parse(ip)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxyUrl)
			tr.ResponseHeaderTimeout = time.Second * time.Duration(5)
		}
	}

	// 创建client
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	// 发请求
	res, err := httpClient.Get(target)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return false

	}

	fmt.Println(ip, " StatusCode : ", res.StatusCode)
	//c, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(" res.body  : ", string(c))
	return true

}
