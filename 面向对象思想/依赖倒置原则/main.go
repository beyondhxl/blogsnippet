package main

import "fmt"

// 游戏
type IPlay interface {
	Play()
}

// 电脑
type PC struct {
}

func (pc *PC) Play() {
	fmt.Print("Play by PC \n")
}

// 手机
type Phone struct {
}

func (ph *Phone) Play() {
	fmt.Print("Play by Phone \n")
}

// 玩家
type Player struct {
}

func (p *Player) PalyZelda(play IPlay) {
	play.Play()
}

func main() {
	pc := &PC{}
	ph := &Phone{}
	player := new(Player)
	player.PalyZelda(pc)
	player.PalyZelda(ph)

	s := &Switch{}
	user := &User{}
	user.PalyZelda(s)
}

type User struct {
}

func (u *User) PalyZelda(s *Switch) {
	s.PalyZelda()
}

type Switch struct {
}

func (s *Switch) PalyZelda() {
	fmt.Print("Play Nintendo Game by Switch")
}
