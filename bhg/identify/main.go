/*
 * @Author: uyplayer
 * @Date: 2022/4/3 22:24
 * @Email: uyplayer@qq.com
 * @File: main.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / bhg/identify
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		fmt.Println(device.Name)
		for _, address := range device.Addresses {
			fmt.Printf("    IP:      %s\n", address.IP)
			fmt.Printf("    Netmask: %s\n", address.Netmask)
		}
	}
}
