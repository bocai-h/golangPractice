package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// singleHandler 代表单次处理函数的类型
type singleHandler func() (data string, n int, err error)

// singleHandler 代表单次处理函数的类型
type handlerConfig struct {
	handler   singleHandler //单次处理函数
	goNum     int           //需要启用的gorouting数量
	number    int           //单个goroutine中的处理次数
	interval  time.Duration //单个goroutine中的处理间隔时间
	counter   int           //数据量计数器，以字节为单位
	counterMu sync.Mutex    //数据量计数器专用的互斥锁
}

// count 会增加计数器的值， 并返回增加后的计数
func (hc *handlerConfig) count(increment int) int {
	hc.counterMu.Lock()
	defer hc.counterMu.Unlock()
	hc.counter += increment
	return hc.counter
}
func main() {
	var mu sync.Mutex
	fmt.Printf("Hello world")
	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			//准备数据
			data = fmt.Sprintf("%s\n", time.Now().Format(time.StampNano))
			//写入数据
			mu.Lock()
			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return
		}
	}

	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			//读取数据
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	var buffer bytes.Buffer
	writingConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}

	readingConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    10,
		number:   2,
		interval: time.Millisecond * 100,
	}

	sign := make(chan struct{}, writingConfig.goNum+readingConfig.goNum)

	//启用多个goroutine对缓冲区进行多次数据写入
	for i := 1; i <= writingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()

			for j := 1; j <= writingConfig.number; j++ {
				time.Sleep(writingConfig.interval)
				data, n, err := writingConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error:%s", i, j, err)
					continue
				}
				total := writingConfig.count(n)
				log.Printf("witer [%d-%d]: %s (total: %d)", i, j, data, total)
			}
		}(i)
	}

	//启用多个goroutine对缓冲区间进行多次数据读取
	for i := 1; i <= readingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()

			for j := 1; j <= readingConfig.number; j++ {
				time.Sleep(readingConfig.interval)
				var data string
				var n int
				var err error
				for {
					data, n, err = readingConfig.handler()
					if err != nil || err != io.EOF {
						break
					}
					//如果读比写快（读时会发生EOF错误）,那就等一会儿
					time.Sleep(readingConfig.interval)
				}
				total := readingConfig.count(n)
				if err != nil {
					log.Printf("reader [%d-%d]:%s (total:%d)", i, j, data, total)
				}
			}
		}(i)
	}
	// signNumber代表需要接受的信号的数量
	signNumber := writingConfig.goNum + readingConfig.goNum
	for j := 0; j < signNumber; j++ {
		<-sign
	}
}
