// 命令模式
// 存在四个角色 客户端、接受者、调用者、命令
// 接受者接收客户端传递来的命令，调用者去调用命令得到结果
// 接受者和调用者互相不知道对方的存在，是一种解藕
package main

import "fmt"

type Command interface {
	Press()
}

type Play struct {
}

func (p *Play) Press() {
	fmt.Println("press play")
}

type Pause struct {
}

func (p *Pause) Press() {
	fmt.Println("press pause")
}

type MusicPlayer struct {
	cmd Command
}

func (m *MusicPlayer) SetCommand(cmd Command) {
	m.cmd = cmd
}

func (m *MusicPlayer) Do() {
	m.cmd.Press()
}
func main() {
	play := &Play{}
	pause := &Pause{}
	mp := &MusicPlayer{}

	mp.cmd = play
	mp.Do()

	mp.cmd = pause
	mp.Do()

}
