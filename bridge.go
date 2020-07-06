// 桥接模式 又称为接口模式
// 这里学习时的例子给的是电脑和打印机的关系
// 为了加深理解我自己用鲜花和颜色重新实现
// 自己的理解：连接两个抽象类，实现两部分自由组合，作为接收方应该有一个可以变换实现的地方，用一个结构体或者map，参数是另一部分的实现
// 参考：https://www.dcdapp.com/article/6787650990091796995
package main

import "fmt"

type Flower interface {
	Show()
	SetColor()
}

type Color interface {
	coloring()
}

type Rose struct {
	color Color
}

func (r *Rose) SetColor(color Color) {
	r.color = color
}

func (r *Rose) Show() {
	fmt.Println("Rose")
	r.color.coloring()
}

type Lily struct {
	color Color
}

func (l *Lily) SetColor(color Color) {
	l.color = color
}

func (l *Lily) Show() {
	fmt.Println("Lily")
	l.color.coloring()
}

type Red struct{}

func (r *Red) coloring() {
	fmt.Println("coloring red")
}

type Pink struct{}

func (p *Pink) coloring() {
	fmt.Println("coloring pink")
}

func main() {
	//实例化
	red := &Red{}
	pink := &Pink{}
	rose := &Rose{}
	lily := &Lily{}

	rose.SetColor(red)
	rose.Show()
	rose.SetColor(pink)
	rose.Show()

	lily.SetColor(red)
	lily.Show()
	lily.SetColor(pink)
	lily.Show()
}
