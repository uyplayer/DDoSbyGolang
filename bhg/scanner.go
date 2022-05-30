/*
 * @Author: uyplayer
 * @Date: 2022/3/28 01:14
 * @Email: uyplayer@qq.com
 * @File: scanner.go
 * @Software: GoLand
 * @Dir: DDoSbyGolang / bhg
 * @Project_Name: DDoSbyGolang
 * @Description:
 */

package bhg

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func SlowScanner() {

	for i := 0; i < 1024; i++ {
		addres := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addres)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf(" %d is open \n")
	}

}

func FastScannerWithGoRoutine() {
	// 该程序无法继续main goroutine 提前结束，
	for i := 0; i < 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

}

func solveproblomabove() {
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)

		}(i)
	}

	wg.Wait()
}

func generator() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// 调用worker
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// 发货
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// 接受结果
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// 要结束
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}

func worker(ports, results chan int) {

	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}

}

func generatorwait() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go workdone(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func workdone(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}
