package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Promise struct {
	wg  sync.WaitGroup
	res string
	err error
}

func NewPromise(f func() (string, error)) *Promise {
	p := &Promise{}
	p.wg.Add(1)
	go func() {
		p.res, p.err = f()
		p.wg.Done()
	}()
	return p
}

func (p *Promise) Then(r func(string), e func(error)) *Promise {
	go func() {
		p.wg.Wait()
		if p.err != nil {
			e(p.err)
			return
		}
		r(p.res)
	}()
	return p
}

func exampleTicker() (string, error) {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		<-time.Tick(time.Second * 1)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(100) % 2
	fmt.Println(r)
	if r != 0 {
		return "hello, world", nil
	} else {
		return "", fmt.Errorf("error")
	}
}

func main() {
	doneChan := make(chan int)

	var p = NewPromise(exampleTicker)
	p.Then(func(result string) { fmt.Println("result: ", result); doneChan <- 1 },
		func(err error) { fmt.Println("err: ", err); doneChan <- 1 })

	<-doneChan
}
