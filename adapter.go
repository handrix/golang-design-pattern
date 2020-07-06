//适配器模式
//适配器模式用于讲一个类的接口转换为业务代码期望的另一个接口
//譬如新的代码要兼容老旧代码的接口
package main

import "fmt"

type VideoPlayer interface {
	play(fileType, name string)
}

// 旧的类
type OlderPlayer struct {
}

func (*OlderPlayer) PlayMP4(name string) {
	fmt.Printf("play %s.mp4\n", name)
}

func (*OlderPlayer) PlayAVI(name string) {
	fmt.Printf("play %s.AVI\n", name)
}

// 适配器类需要持有旧的类
type Adapter struct {
	olderPlayer OlderPlayer
}

func (a *Adapter) play(fileType, name string) {
	switch {
	case fileType == "mp4":
		a.olderPlayer.PlayMP4(name)
	case fileType == "avi":
		a.olderPlayer.PlayAVI(name)
	default:
		fmt.Println("不支持该文件格式")
	}
}

func main() {
	player := Adapter{}
	player.play("mp4", "durarara")
	player.play("avi", "durarara")
	player.play("akb", "durarara")
}
