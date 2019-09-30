package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

// protecting表示是否需要使用互斥锁来保护数据写入
//如果等于0则表示不需要 等于1则表示需要
//多次改变量的值，然后运行程序，观察程序输出
var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1, "It indicates whether to use mutex to protect data writing.")
}
func main() {
	flag.Parse()
	// buffer代表缓存区
	var buffer bytes.Buffer

	const (
		max1 = 5  //代表启用的gorouting的数量
		max2 = 10 //代表每个gorouting需要写入数据块的数量
		max3 = 10 //代表每个数据块中需要有多少个重复的数字
	)

	// 代表以下流程需要使用的互斥锁
	var mu sync.Mutex
	// 代表信号通道
	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 0; j < max2; j++ {
				// 准备数据
				header := fmt.Sprintf("\n[id: %d, iteration: %d]", id, j)
				data := fmt.Sprintf(" %d", id*j)
				// 写入数据
				if protecting > 0 {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error:%s, [%d]", err, id)
				}
				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error:%s, [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}
	for i := 0; i < max1; i++ {
		<-sign
	}
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:%s\n", string(data))
}
