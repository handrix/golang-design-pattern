//组合模式
// 通过操作方法写在抽象接口或者composite处分为
// 透明模式和安全模式
// composite类应有一个数据结构来存放add进来的方法
// 通过composite 完成对不同叶子方法的组合
// 这个模式目前不完整，明天接着搞。。。
package main

import "fmt"

type Graphics interface {
	draw()
}

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("draw a circle")
}

type Rectangle struct {
}

func (r *Rectangle) draw() {
	fmt.Println("draw a rectangle")
}

type Line struct {
}

func (l *Line) draw() {
	fmt.Println("draw a line")
}

type Picture struct {
	vector []Graphics
}

func (p *Picture) draw() {
	fmt.Println("draw a picture...")
	for _, v := range p.vector {
		v.draw()
	}
}

func (p *Picture) add(graphic Graphics) {
	p.vector = append(p.vector, graphic)
}

//func (p *Picture) remove(graphic Graphics)  {
//	delete(p.vector, graphic)
//}

//func (p *Picture) getChild(graphic Graphics) Graphics {
//	return p.vector[graphic]
//}

func main() {
	picture := Picture{}
	line := Line{}
	circle := Circle{}
	rectangle := Rectangle{}

	picture.add(&line)
	picture.add(&line)
	picture.add(&rectangle)
	picture.add(&circle)
	picture.draw()
}
