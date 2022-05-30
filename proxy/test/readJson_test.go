/*
 * @Author: uyplayer
 * @Date: 2022/3/9 11:35 PM
 * @Email: uyplayer@qq.com
 * @File: readJson_test.go
 * @Software: GoLand
 * @Dir: DDoSMethodsbyGolang / proxy
 * @Project_Name: DDoSMethodsbyGolang
 * @Description:
 */

package test

import (
	"DDoSbyGolang/proxy"
	"testing"
)

//
//  TesteadJsonFile
//  @Description: 测试
//  @param t
//
func TestReadJsonFile(t *testing.T) {

	_, err := proxy.ReadJsonFile()
	if err != nil {
		t.Errorf("  %v ", err)
	}
}
