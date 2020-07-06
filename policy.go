// 这个最好理解，根据不同场景、行为选择不同算法
// 避免又臭又长的if else
// 但是即便对算法进行了封装，在实际使用的时候还是避免不了if else判断
// 网上冲浪得到解决方法，初始化时放入map，根据key取用
// https://juejin.im/post/5a31d5de6fb9a0450167f76f
// 但是随后发现弊端，这个作者这种写法只能在同一个函数内完成算法撰写，可读性太差
// 继续冲浪得到了一个外国小哥的解答，比较满意
// https://medium.com/@ricardogomesdnb/simple-strategy-pattern-using-golang-6f287e8914cb
package main

import "fmt"

type Sorter interface {
	Sort()
}

type QS struct{}

func (qs QS) Sort() {
	fmt.Println("quick sort.....")
	fmt.Println("Done")
}

type MS struct{}

func (ms MS) Sort() {
	fmt.Println("merge sort.....")
	fmt.Println("Done")
}

var choice = map[string]Sorter{
	"qs": QS{},
	"ms": MS{},
}

func main() {
	choice["qs"].Sort()
	choice["ms"].Sort()
}
