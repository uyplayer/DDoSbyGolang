/*
 * @Author: uyplayer
 * @Date: 2022/3/9 10:55 PM
 * @Email: uyplayer@qq.com
 * @File: readJson.go
 * @Software: GoLand
 * @Dir: DDoSMethodsbyGolang / proxy
 * @Project_Name: DDoSMethodsbyGolang
 * @Description:
 */

package proxy

import (
	"DDoSbyGolang/common"
	"encoding/json"
	"io/ioutil"
	"os"
)

//
//  JsonDatas
//  @Description: 解释json文件的对象
//
type JsonData struct {
	Proxyroviders []struct {
		Url     string `json:"url"`
		Timeout int    `json:"timeout"`
	} `json:"proxyproviders"`
}

//
//  ReadJsonFile
//  @Description: 解释json文件
//  @return error
//
func ReadJsonFile() (*JsonData, error) {

	// 打开json文件
	openJsonFile, err := os.Open(file)
	defer openJsonFile.Close()
	// 如果打开失败
	if err != nil {
		common.Error.Fatalln(err)
	}

	// 查看文件的状态
	fileInfo, err := openJsonFile.Stat()
	// 如果无法获取文件状态
	if err != nil {
		common.Error.Fatalln(err)
	}
	// 查看文件大小是否空
	if fileInfo.Size() == 0 {
		common.Error.Fatalln(err)
	}

	// 写入
	jsonData, err := ioutil.ReadAll(openJsonFile)
	// 如果无法写入
	if err != nil {
		common.Error.Fatalln(err)
	}
	// 解释
	data := &JsonData{}
	if err := json.Unmarshal(jsonData, data); err != nil {
		common.Error.Fatalln(err)

	}

	return data, nil
}
