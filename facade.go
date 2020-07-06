// 外观模式
// 当一个请求需要系统中多个子服务时，引用该模式
// 即请求只到facade部分，子服务的调用写在facade中
package main

import "fmt"

type Server struct {
	calc Calc
	send Send
}

type Calc struct {
}
type Send struct {
}

func (*Calc) CalcST() {
	fmt.Println("calc something")
}

func (*Send) SendST() {
	fmt.Println("send something")
}

func (s *Server) CalcAndPost() {
	s.calc.CalcST()
	s.send.SendST()
}

func main() {
	s := &Server{calc: Calc{}, send: Send{}}
	s.CalcAndPost()
}
