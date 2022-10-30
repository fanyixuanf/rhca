/*
@Time : 2022/10/30 15:45
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : mediator_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"strings"
	"testing"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadDate() {
	c.Data = "music,image"
	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	//
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]
	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound,c.Video)
	//
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	//
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard)Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	//
	GetMediatorInstance().changed(s)
}

type Mediator struct {
	CD *CDDriver
	CPU *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}

func TestMediator (t *testing.T) {
	med := GetMediatorInstance()
	med.CD = &CDDriver{}
	med.CPU = &CPU{}
	med.Video = &VideoCard{}
	med.Sound = &SoundCard{}

	med.CD.ReadDate()
	fmt.Println("-----------")
	fmt.Printf("CD: %s\n", med.CD.Data)
	fmt.Printf("CPU Sound:  %s\n", med.CPU.Sound)
	fmt.Printf("CPU Video:  %s\n", med.CPU.Video)
	fmt.Printf("Video:  %s\n", med.Video.Data)
	fmt.Printf("Sound:  %s\n", med.Sound.Data)
}