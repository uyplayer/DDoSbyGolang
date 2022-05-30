/*
 * @Author: uyplayer
 * @Date: 2022/3/11 7:53 PM
 * @Email: uyplayer@qq.com
 * @File: errors.go
 * @Software: GoLand
 * @Dir: DDoS Attack Script With 44 Methods by Golang /
 * @Project_Name: DDoS Attack Script With 44 Methods by Golang
 * @Description:
 */

package common

type UserError struct {
	createTime string
	message    string
}

func (ue *UserError) PrintError() {

}
