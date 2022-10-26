/*
@Time : 2022/10/25 20:35
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : bridge_test.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplement interface {
	Send(text, to string)
}


type MessageSMS struct {}

func VisSMS() MessageImplement {
	return &MessageSMS{}
}

func (sms *MessageSMS)Send(text, to string) {
	fmt.Printf("Send %s to %s via sms\n", text, to)
}

type MessageEmail struct {}

func ViaEmail() MessageImplement {
	return &MessageEmail{}
}

func (e *MessageEmail)Send (text, to string) {
	fmt.Printf("Send %s to %s mail \n", text, to)
}

type CommonMessage struct {
	method MessageImplement
}

func NewCommonMessage(method MessageImplement) *CommonMessage {
	return &CommonMessage{method: method}
}

func (c *CommonMessage)SendMessage(text, to string) {
	c.method.Send(text ,to)
}

type UrgencyMessage struct {
	method MessageImplement
}

func NewUrgencyMessage(method MessageImplement) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (u *UrgencyMessage)SendMessage(text, to string) {
	u.method.Send(text, to)
}


func TestMessage(t *testing.T) {
	cms := NewCommonMessage(VisSMS())
	cms.SendMessage("吃了吗", "zhangsan")

	cml := NewCommonMessage(ViaEmail())
	cml.SendMessage("吃了吗", "zhangsan")

	ums := NewUrgencyMessage(VisSMS())
	ums.SendMessage("吃了吗", "zhangsan")

	uml := NewUrgencyMessage(ViaEmail())
	uml.SendMessage("吃了吗", "zhangsan")

}