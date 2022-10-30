/*
@Time : 2022/10/30 16:21
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : memento_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Memento interface {}

type Game struct {
	hp,
	mp int
}

type gameMemento struct {
	hp,
	mp int
}

func (g *Game) Play(mp, hp int) {
	g.mp += mp
	g.hp += hp
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load (m Memento) {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp
}

func (g *Game) Status () {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}

func TestMemento(t *testing.T) {
	game := &Game{
		hp: 10,
		mp: 10,
	}
	game.Status()
	progress := game.Save()

	game.Play(-1, -1)
	game.Status()

	game.Load(progress)
	game.Status()
}