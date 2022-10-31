/*
@Time : 2022/10/31 20:21
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : strategy_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type PaymentContext struct {
	Name, CardID string
	Money int
}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Payment struct {
	context *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context:  &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type Cash struct {}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct {}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}

func TestStrategy(t *testing.T) {

	cash := NewPayment("CET4", "", 1200,&Cash{})
	cash.Pay()

	bank := NewPayment("CET6", "cet6666", 1500, &Bank{})
	bank.Pay()

}