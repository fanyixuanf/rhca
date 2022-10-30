/*
@Time : 2022/10/30 14:15
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : command_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

type MotherBoard struct {}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb: mb}
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

func (*MotherBoard)Start() {
	fmt.Println("system start")
}

func (*MotherBoard)Reboot () {
	fmt.Println("system reboot")
}

type RebootCommand struct {
	mb *MotherBoard
}

func (r *RebootCommand) Execute() {
	r.mb.Reboot()
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb: mb}
}

type Box struct {
	btn1 Command
	btn2 Command
}

func (b *Box) PressBtn1() {
	b.btn1.Execute()
}

func (b *Box) PressBtn2() {
	b.btn2.Execute()
}

func NewBox(b1, b2 Command) *Box {
	return &Box{btn1: b1, btn2: b2}
}


func TestCommand(t *testing.T) {
	mb := &MotherBoard{}
	sc := NewStartCommand(mb)
	rc := NewRebootCommand(mb)

	b1 := NewBox(sc, rc)
	b1.PressBtn1()
	b1.PressBtn2()

	b2 := NewBox(sc, rc)
	b2.PressBtn1()
	b2.PressBtn2()

}