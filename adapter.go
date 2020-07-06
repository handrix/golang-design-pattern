package main

import "fmt"

type VideoPlayer interface {
	play(fileType, name string)
}

type OlderPlayer struct {
}

func (*OlderPlayer) PlayMP4(name string) {
	fmt.Printf("play %s.mp4\n", name)
}

func (*OlderPlayer) PlayAVI(name string) {
	fmt.Printf("play %s.AVI\n", name)
}

// 持有旧接口
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
